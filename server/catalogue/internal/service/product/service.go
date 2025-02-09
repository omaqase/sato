package product

import (
	"context"
	"github.com/omaqase/sato/catalogue/internal/repository/product"
	"github.com/omaqase/sato/catalogue/internal/service/jwt"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/catalogue"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Service struct {
	repository product.IRepository
	jwtService jwt.IService
}

func NewService(repository product.IRepository, jwtService jwt.IService) protobuf.ProductServiceServer {
	return &Service{
		repository: repository,
		jwtService: jwtService,
	}
}

func (s Service) CreateProduct(ctx context.Context, request *protobuf.CreateProductRequest) (*protobuf.CreateProductResponse, error) {
	if err := ValidateCreateProductRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.jwtService.Parse(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Println(userID)

	return nil, nil
}

func (s Service) UpdateProduct(ctx context.Context, request *protobuf.UpdateProductRequest) (*protobuf.UpdateProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetProductByID(ctx context.Context, request *protobuf.GetProductByIDRequest) (*protobuf.GetProductByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListProduct(ctx context.Context, request *protobuf.ListProductRequest) (*protobuf.ListProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteProduct(ctx context.Context, request *protobuf.DeleteProductRequest) (*protobuf.DeleteProductResponse, error) {
	//TODO implement me
	panic("implement me")
}
