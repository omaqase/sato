package product

import (
	"errors"
	"github.com/google/uuid"
	protobuf "github.com/omaqase/sato/product/pkg/api/product/user"
)

func ValidateCreateProductRequest(request *protobuf.CreateProductRequest) error {
	if request.Title == "" || request.Price <= 0 {
		return errors.New("invalid product data: title and price are required")
	}
	if _, err := uuid.Parse(request.CategoryId); err != nil {
		return errors.New("invalid category ID format")
	}
	return nil
}

func ValidateGetProductByIdRequest(request *protobuf.GetProductByIdRequest) error {
	if _, err := uuid.Parse(request.Id); err != nil {
		return errors.New("invalid product ID format")
	}
	return nil
}

func ValidateUpdateProductRequest(request *protobuf.UpdateProductRequest) error {
	if _, err := uuid.Parse(request.Id); err != nil {
		return errors.New("invalid product ID format")
	}
	if request.Title == "" || request.Price <= 0 {
		return errors.New("invalid product data: title and price are required")
	}
	if _, err := uuid.Parse(request.CategoryId); err != nil {
		return errors.New("invalid category ID format")
	}
	return nil
}

func ValidateDeleteProductRequest(request *protobuf.DeleteProductRequest) error {
	if _, err := uuid.Parse(request.Id); err != nil {
		return errors.New("invalid product ID format")
	}
	return nil
}

func ValidateListProductsRequest(request *protobuf.ListProductsRequest) error {
	if request.Page <= 0 || request.Limit <= 0 || request.Limit > 100 {
		return errors.New("page and limit must be positive integers, and limit must not exceed 100")
	}
	return nil
}
