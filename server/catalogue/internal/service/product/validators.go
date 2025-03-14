package product

import (
	"errors"
	"fmt"
	"github.com/omaqase/sato/catalogue/internal/service"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/catalogue"
	"slices"
)

var ApproveStatus = []string{"pending", "approved", "rejected"}

func ValidateApproveStatus(status string) error {
	if !slices.Contains(ApproveStatus, status) {
		return fmt.Errorf("invalid approve status")
	}

	return nil
}

//func ValidateCreateProductRequest(request *protobuf.CreateProductRequest) error {
//	if request == nil {
//		return errors.New("request is nil")
//	}
//
//	if len(request.Title) == 0 {
//		return errors.New("title must not be empty")
//	}
//
//	if err := service.ValidateUUID(request.CategoryId); err != nil {
//		return errors.New("invalid category id")
//	}
//
//	if err := service.ValidateUUID(request.SellerId); err != nil {
//		return errors.New("invalid seller id")
//	}
//
//	if request.Stock < 0 {
//		return errors.New("stock must not be negative")
//	}
//
//	return nil
//}
//
//func ValidateUpdateProductRequest(request *protobuf.UpdateProductRequest) error {
//	if request == nil {
//		return errors.New("request is nil")
//	}
//
//	if err := service.ValidateUUID(request.ProductId); err != nil {
//		return errors.New("invalid product id")
//	}
//
//	if request.Title != nil && len(*request.Title) == 0 {
//		return errors.New("title must not be empty")
//	}
//
//	if request.CategoryId != nil {
//		if err := service.ValidateUUID(*request.CategoryId); err != nil {
//			return errors.New("invalid category id")
//		}
//	}
//
//	if request.Stock != nil && *request.Stock < 0 {
//		return errors.New("stock must not be negative")
//	}
//
//	return nil
//}
//
//func ValidateListProductsRequest(request *protobuf.ListProductsRequest) error {
//	if request == nil {
//		return errors.New("request is nil")
//	}
//
//	if request.Limit <= 0 || request.Limit > 100 {
//		return errors.New("limit must be between 1 and 100")
//	}
//
//	if request.Offset < 0 {
//		return errors.New("offset must not be negative")
//	}
//
//	return nil
//}

func ValidateCreateProductRequest(request *protobuf.CreateProductRequest) error {
	if request == nil {
		return errors.New("request is nil")
	}

	if request.Stock < 0 {
		return errors.New("stock is negative")
	}

	if err := service.ValidateUUID(request.CategoryId); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateProductRequest(request *protobuf.UpdateProductRequest) error {
	if request.ProductId == "" {
		return fmt.Errorf("product_id is required")
	}
	if request.Title == nil || *request.Title == "" {
		return fmt.Errorf("title is required")
	}
	if request.Stock == nil || *request.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return nil
}
