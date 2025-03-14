package product

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/omaqase/sato/catalogue/internal/repository/favorites"

	"github.com/jackc/pgx/v5"
	"github.com/omaqase/sato/catalogue/internal/repository/product"
	"github.com/omaqase/sato/catalogue/internal/service"
	"github.com/omaqase/sato/catalogue/internal/service/jwt"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/catalogue"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repository    product.IRepository
	favoritesRepo favorites.Repository
	jwtService    jwt.IService
}

func (s Service) GetProductByID(ctx context.Context, request *protobuf.GetProductByIDRequest) (*protobuf.GetProductByIDResponse, error) {
	if err := service.ValidateUUID(request.ProductId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	productEntity, err := s.repository.GetByID(ctx, request.ProductId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &protobuf.GetProductByIDResponse{
		ProductId:  productEntity.ID,
		CategoryId: productEntity.CategoryID,
		SellerId:   productEntity.SellerID,
		Title:      productEntity.Title,
		Stock:      int32(productEntity.Stock),
		Rating:     float32(productEntity.Rating),
		Approved:   true,
		CreatedAt:  productEntity.CreatedAt.Format(time.DateTime),
		UpdatedAt:  productEntity.UpdatedAt.Format(time.DateTime),
	}

	return response, nil
}

func (s Service) ListProduct(ctx context.Context, request *protobuf.ListProductRequest) (*protobuf.ListProductResponse, error) {
	// Проверка входных параметров
	limit := int(request.Limit)
	offset := int(request.Offset)
	if limit <= 0 || offset < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid limit or offset")
	}

	// Получение списка продуктов из репозитория
	products, err := s.repository.List(ctx, limit, offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Преобразование сущностей домена в структуры Protobuf
	var productList []*protobuf.ListProductItem
	for _, productEntity := range products {
		productList = append(productList, &protobuf.ListProductItem{
			ProductId:  productEntity.ID,
			CategoryId: productEntity.CategoryID,
			SellerId:   productEntity.SellerID,
			Title:      productEntity.Title,
			Stock:      int32(productEntity.Stock),
			Rating:     float32(productEntity.Rating),
			Approved:   true,
			CreatedAt:  productEntity.CreatedAt.Format(time.DateTime),
			UpdatedAt:  productEntity.UpdatedAt.Format(time.DateTime),
		})
	}

	// Формирование ответа
	response := &protobuf.ListProductResponse{
		Products: productList,
	}

	return response, nil
}

func NewService(repository product.IRepository, jwtService jwt.IService) protobuf.ProductServiceServer {
	return &Service{repository: repository, jwtService: jwtService}
}

func (s Service) CreateProduct(ctx context.Context, request *protobuf.CreateProductRequest) (*protobuf.CreateProductResponse, error) {
	log.Printf("request: %v", request)

	if err := ValidateCreateProductRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sellerId, err := s.jwtService.Parse(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("sellerId: %v", sellerId)

	contract := product.NewCreateProductContract(sellerId, request.Title, request.Slug, request.CategoryId, int(request.Stock))

	dest, err := s.repository.Create(ctx, contract)
	if err != nil {
		log.Printf("err: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &protobuf.CreateProductResponse{
		ProductId:     dest.ID,
		CategoryId:    dest.CategoryID,
		SellerId:      dest.SellerID,
		Title:         dest.Title,
		Stock:         int32(dest.Stock),
		Rating:        float32(dest.Rating),
		ApproveStatus: dest.Approved,
		CreatedAt:     dest.CreatedAt.Format(time.DateTime),
	}

	return out, nil
}

func (s Service) UpdateProduct(ctx context.Context, request *protobuf.UpdateProductRequest) (*protobuf.UpdateProductResponse, error) {
	if err := ValidateUpdateProductRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sellerId, err := s.jwtService.Parse(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to parse JWT token")
	}

	contract := product.NewUpdateProductContract(
		request.ProductId,
		*request.Title,
		"slug",
		sellerId,
		int(*request.Stock),
		sellerId,
	)

	dest, err := s.repository.Update(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &protobuf.UpdateProductResponse{
		ProductId: dest.ID,
		Title:     dest.Title,
		Stock:     int32(dest.Stock),
		UpdatedAt: dest.UpdatedAt.Format(time.DateTime),
	}

	return response, nil
}

func (s Service) DeleteProduct(ctx context.Context, request *protobuf.DeleteProductRequest) (*protobuf.DeleteProductResponse, error) {
	if err := service.ValidateUUID(request.ProductId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.repository.DeleteByID(ctx, request.ProductId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.DeleteProductResponse{}, nil
}

func (s Service) GetProductBySlug(ctx context.Context, req *protobuf.GetProductBySlugRequest) (*protobuf.GetProductBySlugResponse, error) {
	product, err := s.repository.GetBySlug(ctx, req.Slug)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product by slug: %v", err)
	}

	return &protobuf.GetProductBySlugResponse{
		ProductId:  product.ID,
		CategoryId: product.CategoryID,
		SellerId:   product.SellerID,
		Title:      product.Title,
		Stock:      int32(product.Stock),
		Rating:     float32(product.Rating),
		Approved:   product.Approved == "approved",
		CreatedAt:  product.CreatedAt.Format(time.DateTime),
		UpdatedAt:  product.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (s *Service) AddToFavorites(ctx context.Context, request *protobuf.AddToFavoritesRequest) (*protobuf.AddToFavoritesResponse, error) {
	err := s.favoritesRepo.AddToFavorites(ctx, request.UserId, request.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.AddToFavoritesResponse{}, nil
}

func (s *Service) RemoveFromFavorites(ctx context.Context, request *protobuf.RemoveFromFavoritesRequest) (*protobuf.RemoveFromFavoritesResponse, error) {
	err := s.favoritesRepo.RemoveFromFavorites(ctx, request.UserId, request.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.RemoveFromFavoritesResponse{}, nil
}

func (s *Service) GetUserFavorites(ctx context.Context, request *protobuf.GetUserFavoritesRequest) (*protobuf.GetUserFavoritesResponse, error) {
	favorites, err := s.favoritesRepo.GetUserFavorites(ctx, request.UserId, int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user favorites: %v", err)
	}

	out := &protobuf.GetUserFavoritesResponse{
		ProductIds: favorites,
	}

	return out, nil
}
