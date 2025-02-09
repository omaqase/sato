package category

import (
	"errors"
	"github.com/omaqase/sato/catalogue/internal/service"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/catalogue"
)

func ValidateCreateCategoryRequest(request *protobuf.CreateCategoryRequest) error {
	if request == nil {
		return errors.New("request is nil")
	}

	if len(request.Title) <= 0 {
		return errors.New("title must be greater than zero")
	}

	if len(request.Slug) <= 0 {
		return errors.New("slug must be greater than zero")
	}

	if err := service.ValidateSlug(request.Slug); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateCategoryRequest(request *protobuf.UpdateCategoryRequest) error {
	if request == nil {
		return errors.New("request is nil")
	}

	if err := service.ValidateUUID(request.CategoryId); err != nil {
		return err
	}

	if len(*request.Title) <= 0 {
		return errors.New("title must be greater than zero")
	}

	if len(*request.Slug) <= 0 {
		return errors.New("slug must be greater than zero")
	}

	if err := service.ValidateSlug(*request.Slug); err != nil {
		return err
	}

	return nil
}

func ValidateListCategoryRequest(request *protobuf.ListCategoryRequest) error {
	if request == nil {
		return errors.New("request is nil")
	}

	if request.Limit <= 0 || request.Limit > 100 {
		return errors.New("limit must be between 0 and 100")
	}

	return nil
}
