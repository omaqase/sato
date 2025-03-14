package category

import (
	"errors"
	"log"

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

	log.Println(request.Limit)
	log.Println(request.Offset)

	if request.Limit <= 0 || request.Limit > 100 {
		return errors.New("limit must be greater than 0 and less than or equal to 100")
	}

	if request.Offset < 0 {
		return errors.New("offset must be greater than or equal to 0")
	}

	return nil
}
