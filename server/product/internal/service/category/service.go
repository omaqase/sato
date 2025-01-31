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

// CreateCategory creates a new category.
func (c *CategoryService) CreateCategory(ctx context.Context, request *protobuf.CreateCategoryRequest) (*protobuf.CreateCategoryResponse, error) {
	if err := ValidateCreateCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := category.NewCreateCategoryContract(request.Title, request.Slug)
	category, err := c.repository.CreateCategory(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrInvalidSlug) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.CreateCategoryResponse{
		Id:    category.ID.String(),
		Title: category.Title,
		Slug:  category.Slug,
	}, nil
}

// GetCategoryById retrieves a category by its ID.
func (c *CategoryService) GetCategoryById(ctx context.Context, request *protobuf.GetCategoryByIdRequest) (*protobuf.GetCategoryByIdResponse, error) {
	if err := ValidateGetCategoryByIdRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	categoryID, _ := uuid.Parse(request.Id)
	category, err := c.repository.GetCategoryByID(ctx, categoryID)
	if err != nil {
		if errors.Is(err, postgres.ErrCategoryNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.GetCategoryByIdResponse{
		Id:    category.ID.String(),
		Title: category.Title,
		Slug:  category.Slug,
	}, nil
}

// UpdateCategory updates an existing category.
func (c *CategoryService) UpdateCategory(ctx context.Context, request *protobuf.UpdateCategoryRequest) (*protobuf.UpdateCategoryResponse, error) {
	if err := ValidateUpdateCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	categoryID, _ := uuid.Parse(request.Id)
	contract := category.NewUpdateCategoryContract(categoryID, request.Title, request.Slug)

	category, err := c.repository.UpdateCategory(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrCategoryNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if errors.Is(err, postgres.ErrInvalidSlug) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.UpdateCategoryResponse{
		Id:    category.ID.String(),
		Title: category.Title,
		Slug:  category.Slug,
	}, nil
}

// DeleteCategory deletes a category by its ID.
func (c *CategoryService) DeleteCategory(ctx context.Context, request *protobuf.DeleteCategoryRequest) (*protobuf.DeleteCategoryResponse, error) {
	// Validate the request
	if err := ValidateDeleteCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var err error

	// Handle the oneof Identifier field
	switch identifier := request.Identifier.(type) {
	case *protobuf.DeleteCategoryRequest_Id:
		// Parse the ID and delete by ID
		categoryID, parseErr := uuid.Parse(identifier.Id)
		if parseErr != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid category ID format")
		}
		err = c.repository.DeleteCategory(ctx, category.NewDeleteCategoryContractByID(categoryID))

	case *protobuf.DeleteCategoryRequest_Slug:
		// Delete by Slug
		err = c.repository.DeleteCategory(ctx, category.NewDeleteCategoryContractBySlug(identifier.Slug))

	default:
		// This case should never happen due to validation, but handle it just in case
		return nil, status.Error(codes.InvalidArgument, "invalid identifier type")
	}

	// Handle repository errors
	if err != nil {
		if errors.Is(err, postgres.ErrCategoryNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return success response
	return &protobuf.DeleteCategoryResponse{
		Status: 0, // Success
	}, nil
}

// ListCategories retrieves a paginated list of categories.
func (c *CategoryService) ListCategories(ctx context.Context, request *protobuf.ListCategoriesRequest) (*protobuf.ListCategoriesResponse, error) {
	if err := ValidateListCategoriesRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := category.NewListCategoriesContract(int(request.Page), int(request.Limit))
	categories, totalCount, err := c.repository.ListCategories(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var responseItems []*protobuf.GetCategoryByIdResponse
	for _, category := range categories {
		responseItems = append(responseItems, &protobuf.GetCategoryByIdResponse{
			Id:    category.ID.String(),
			Title: category.Title,
			Slug:  category.Slug,
		})
	}

	return &protobuf.ListCategoriesResponse{
		Categories: responseItems,
		TotalCount: int32(totalCount),
	}, nil
}
