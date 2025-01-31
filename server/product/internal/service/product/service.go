package product

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/omaqase/sato/product/internal/repository/postgres"
	"github.com/omaqase/sato/product/internal/repository/postgres/product"
	protobuf "github.com/omaqase/sato/product/pkg/api/product/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ProductService struct {
	repository product.IProductRepository
}

func NewProductService(repository product.IProductRepository) protobuf.ProductServiceServer {
	return &ProductService{repository: repository}
}

// CreateProduct creates a new product.
func (p *ProductService) CreateProduct(ctx context.Context, request *protobuf.CreateProductRequest) (*protobuf.CreateProductResponse, error) {
	if err := ValidateCreateProductRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	categoryID, _ := uuid.Parse(request.CategoryId)
	contract := product.NewCreateProductContract(
		request.Title,
		request.Description,
		request.Price,
		categoryID,
	)

	product, err := p.repository.CreateProduct(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.CreateProductResponse{
		Id:          product.ID.String(),
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		CategoryId:  product.CategoryID.String(),
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (p *ProductService) GetProductById(ctx context.Context, request *protobuf.GetProductByIdRequest) (*protobuf.GetProductByIdResponse, error) {
	if err := ValidateGetProductByIdRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	productID, _ := uuid.Parse(request.Id)
	product, err := p.repository.GetProductByID(ctx, productID)
	if err != nil {
		if errors.Is(err, postgres.ErrProductNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.GetProductByIdResponse{
		Id:          product.ID.String(),
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		CategoryId:  product.CategoryID.String(),
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, request *protobuf.UpdateProductRequest) (*protobuf.UpdateProductResponse, error) {
	if err := ValidateUpdateProductRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	productID, _ := uuid.Parse(request.Id)
	categoryID, _ := uuid.Parse(request.CategoryId)
	contract := product.NewUpdateProductContract(
		productID,
		request.Title,
		request.Description,
		request.Price,
		categoryID,
	)

	product, err := p.repository.UpdateProduct(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrProductNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.UpdateProductResponse{
		Id:          product.ID.String(),
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		CategoryId:  product.CategoryID.String(),
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, request *protobuf.DeleteProductRequest) (*protobuf.DeleteProductResponse, error) {
	if err := ValidateDeleteProductRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	productID, _ := uuid.Parse(request.Id)
	err := p.repository.DeleteProduct(ctx, product.NewDeleteProductContract(productID))
	if err != nil {
		if errors.Is(err, postgres.ErrProductNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.DeleteProductResponse{
		Status: 0,
	}, nil
}

func (p *ProductService) ListProducts(ctx context.Context, request *protobuf.ListProductsRequest) (*protobuf.ListProductsResponse, error) {
	if err := ValidateListProductsRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := product.NewListProductsContract(int(request.Page), int(request.Limit))
	products, totalCount, err := p.repository.ListProducts(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var responseItems []*protobuf.GetProductByIdResponse
	for _, product := range products {
		responseItems = append(responseItems, &protobuf.GetProductByIdResponse{
			Id:          product.ID.String(),
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			CategoryId:  product.CategoryID.String(),
			CreatedAt:   product.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &protobuf.ListProductsResponse{
		Products:   responseItems,
		TotalCount: int32(totalCount),
	}, nil
}
