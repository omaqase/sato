package category

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/omaqase/sato/product/internal/repository/postgres"
	"github.com/omaqase/sato/product/internal/repository/postgres/category"
	protobuf "github.com/omaqase/sato/product/pkg/api/product/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	repository category.ICategoryRepository
}

func NewCategoryService(repository category.ICategoryRepository) protobuf.CategoryServiceServer {
	return &CategoryService{repository: repository}
}

func (c *CategoryService) CreateCategory(ctx context.Context, request *protobuf.CreateCategoryRequest) (*protobuf.CreateCategoryResponse, error) {
	if err := ValidateCreateCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	contract := category.NewCreateCategoryContract(request.Title, request.Slug)
	category, err := c.repository.CreateCategory(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrInvalidSlug) {
			return nil, status.Error(codes.AlreadyExists, "slug already exists")
		}
		return nil, status.Error(codes.Internal, "failed to create category")
	}
	return &protobuf.CreateCategoryResponse{
		Id:    category.ID.String(),
		Title: category.Title,
		Slug:  category.Slug,
	}, nil
}

func (c *CategoryService) GetCategoryById(ctx context.Context, request *protobuf.GetCategoryByIdRequest) (*protobuf.GetCategoryByIdResponse, error) {
	if err := ValidateGetCategoryByIdRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	categoryID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid category ID format")
	}
	category, err := c.repository.GetCategoryByID(ctx, categoryID)
	if err != nil {
		if errors.Is(err, postgres.ErrCategoryNotFound) {
			return nil, status.Error(codes.NotFound, "category not found")
		}
		return nil, status.Error(codes.Internal, "failed to retrieve category")
	}
	return &protobuf.GetCategoryByIdResponse{
		Id:    category.ID.String(),
		Title: category.Title,
		Slug:  category.Slug,
	}, nil
}

func (c *CategoryService) UpdateCategory(ctx context.Context, request *protobuf.UpdateCategoryRequest) (*protobuf.UpdateCategoryResponse, error) {
	if err := ValidateUpdateCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	categoryID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid category ID format")
	}
	contract := category.NewUpdateCategoryContract(categoryID, request.Title, request.Slug)
	category, err := c.repository.UpdateCategory(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrCategoryNotFound) {
			return nil, status.Error(codes.NotFound, "category not found")
		}
		if errors.Is(err, postgres.ErrInvalidSlug) {
			return nil, status.Error(codes.AlreadyExists, "slug already exists")
		}
		return nil, status.Error(codes.Internal, "failed to update category")
	}
	return &protobuf.UpdateCategoryResponse{
		Id:    category.ID.String(),
		Title: category.Title,
		Slug:  category.Slug,
	}, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, request *protobuf.DeleteCategoryRequest) (*protobuf.DeleteCategoryResponse, error) {
	if err := ValidateDeleteCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var err error
	switch identifier := request.Identifier.(type) {
	case *protobuf.DeleteCategoryRequest_Id:
		categoryID, parseErr := uuid.Parse(identifier.Id)
		if parseErr != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid category ID format")
		}
		err = c.repository.DeleteCategory(ctx, category.NewDeleteCategoryContractByID(categoryID))
	case *protobuf.DeleteCategoryRequest_Slug:
		err = c.repository.DeleteCategory(ctx, category.NewDeleteCategoryContractBySlug(identifier.Slug))
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid identifier type")
	}
	if err != nil {
		if errors.Is(err, postgres.ErrCategoryNotFound) {
			return nil, status.Error(codes.NotFound, "category not found")
		}
		return nil, status.Error(codes.Internal, "failed to delete category")
	}
	return &protobuf.DeleteCategoryResponse{Status: 0}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, request *protobuf.ListCategoriesRequest) (*protobuf.ListCategoriesResponse, error) {
	if err := ValidateListCategoriesRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	contract := category.NewListCategoriesContract(int(request.Page), int(request.Limit))
	categories, totalCount, err := c.repository.ListCategories(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list categories")
	}
	responseItems := make([]*protobuf.GetCategoryByIdResponse, len(categories))
	for i, category := range categories {
		responseItems[i] = &protobuf.GetCategoryByIdResponse{
			Id:    category.ID.String(),
			Title: category.Title,
			Slug:  category.Slug,
		}
	}
	return &protobuf.ListCategoriesResponse{
		Categories: responseItems,
		TotalCount: int32(totalCount),
	}, nil
}
