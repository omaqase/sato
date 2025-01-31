package category

import (
	"errors"
	protobuf "github.com/omaqase/sato/product/pkg/api/product/user"
)

func ValidateCreateCategoryRequest(request *protobuf.CreateCategoryRequest) error {
	if request.Title == "" || request.Slug == "" {
		return errors.New("title and slug are required")
	}
	return nil
}

func ValidateGetCategoryByIdRequest(request *protobuf.GetCategoryByIdRequest) error {
	if request.Id == "" {
		return errors.New("category ID is required")
	}
	return nil
}

func ValidateUpdateCategoryRequest(request *protobuf.UpdateCategoryRequest) error {
	if request.Id == "" || request.Title == "" || request.Slug == "" {
		return errors.New("ID, title, and slug are required")
	}
	return nil
}

func ValidateDeleteCategoryRequest(request *protobuf.DeleteCategoryRequest) error {
	// Check if neither ID nor Slug is provided
	if request.Identifier == nil {
		return errors.New("either category ID or slug is required")
	}

	switch identifier := request.Identifier.(type) {
	case *protobuf.DeleteCategoryRequest_Id:
		if identifier.Id == "" {
			return errors.New("category ID cannot be empty")
		}
	case *protobuf.DeleteCategoryRequest_Slug:
		if identifier.Slug == "" {
			return errors.New("category slug cannot be empty")
		}
	default:
		return errors.New("invalid identifier type")
	}

	return nil
}

func ValidateListCategoriesRequest(request *protobuf.ListCategoriesRequest) error {
	if request.Page <= 0 || request.Limit <= 0 || request.Limit > 100 {
		return errors.New("page and limit must be positive integers, and limit must not exceed 100")
	}
	return nil
}
