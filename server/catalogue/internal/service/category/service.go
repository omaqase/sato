package category

import (
	"context"
	"log"
	"time"

	"github.com/omaqase/sato/catalogue/internal/repository/category"
	"github.com/omaqase/sato/catalogue/internal/service"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/catalogue"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repository category.IRepository
}

func NewService(repository category.IRepository) protobuf.CategoryServiceServer {
	return &Service{
		repository: repository,
	}
}

func (s Service) CreateCategory(ctx context.Context, request *protobuf.CreateCategoryRequest) (*protobuf.CreateCategoryResponse, error) {
	if err := ValidateCreateCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := category.NewCreateCategoryContract(request.Title, request.Slug)

	dest, err := s.repository.Create(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &protobuf.CreateCategoryResponse{
		Title:     dest.Title,
		Slug:      dest.Slug,
		CreatedAt: dest.CreatedAt.Format(time.DateTime),
	}

	return out, nil
}

func (s Service) UpdateCategory(ctx context.Context, request *protobuf.UpdateCategoryRequest) (*protobuf.UpdateCategoryResponse, error) {
	log.Printf("request: %v", request)

	if err := ValidateUpdateCategoryRequest(request); err != nil {
		log.Printf("err: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := category.NewUpdateCategoryContract(request.CategoryId, *request.Title, *request.Slug)

	dest, err := s.repository.Update(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &protobuf.UpdateCategoryResponse{
		CategoryId: dest.ID,
		Title:      dest.Title,
		Slug:       dest.Slug,
		UpdatedAt:  dest.UpdatedAt.Format(time.DateTime),
	}

	return out, nil
}

func (s Service) GetCategoryByID(ctx context.Context, request *protobuf.GetCategoryByIDRequest) (*protobuf.GetCategoryByIDResponse, error) {
	if err := service.ValidateUUID(request.CategoryId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dest, err := s.repository.GetByID(ctx, request.CategoryId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &protobuf.GetCategoryByIDResponse{
		CategoryId: dest.ID,
		Title:      dest.Title,
		Slug:       dest.Slug,
		CreatedAt:  dest.CreatedAt.Format(time.DateTime),
		UpdatedAt:  dest.UpdatedAt.Format(time.DateTime),
	}

	return out, nil
}

func (s Service) GetCategoryBySlug(ctx context.Context, request *protobuf.GetCategoryBySlugRequest) (*protobuf.GetCategoryBySlugResponse, error) {
	if err := service.ValidateSlug(request.Slug); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dest, err := s.repository.GetBySlug(ctx, request.Slug)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &protobuf.GetCategoryBySlugResponse{
		CategoryId: dest.ID,
		Title:      dest.Title,
		Slug:       dest.Slug,
		CreatedAt:  dest.CreatedAt.Format(time.DateTime),
		UpdatedAt:  dest.UpdatedAt.Format(time.DateTime),
	}

	return out, nil
}

func (s Service) ListCategory(ctx context.Context, request *protobuf.ListCategoryRequest) (*protobuf.ListCategoryResponse, error) {
	if err := ValidateListCategoryRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dest, err := s.repository.List(ctx, int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var categories []*protobuf.ListCategoryItem

	for _, item := range dest {
		categoryItem := &protobuf.ListCategoryItem{
			CategoryId: item.ID,
			Title:      item.Title,
			Slug:       item.Slug,
			CreatedAt:  item.CreatedAt.Format(time.DateTime),
			UpdatedAt:  item.UpdatedAt.Format(time.DateTime),
		}

		categories = append(categories, categoryItem)
	}

	out := &protobuf.ListCategoryResponse{
		Categories: categories,
	}

	return out, nil
}

func (s Service) DeleteCategoryByID(ctx context.Context, request *protobuf.DeleteCategoryByIDRequest) (*protobuf.DeleteCategoryByIDResponse, error) {
	if err := service.ValidateUUID(request.CategoryId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.repository.DeleteByID(ctx, request.CategoryId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.DeleteCategoryByIDResponse{}, nil
}

func (s Service) DeleteCategoryBySlug(ctx context.Context, request *protobuf.DeleteCategoryBySlugRequest) (*protobuf.DeleteCategoryBySlugResponse, error) {
	if err := service.ValidateSlug(request.Slug); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.repository.DeleteBySlug(ctx, request.Slug); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.DeleteCategoryBySlugResponse{}, nil
}

func (s Service) GetCategoryAncestors(ctx context.Context, req *protobuf.GetCategoryAncestorsRequest) (*protobuf.GetCategoryAncestorsResponse, error) {
	ancestors, err := s.repository.GetAncestors(ctx, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get category ancestors: %v", err)
	}

	var protoAncestors []*protobuf.GetCategoryAncestorsItem
	for _, ancestor := range ancestors {
		protoAncestors = append(protoAncestors, &protobuf.GetCategoryAncestorsItem{
			CategoryId: ancestor.ID,
			Title:      ancestor.Title,
			Slug:       ancestor.Slug,
			CreatedAt:  ancestor.CreatedAt.Format(time.DateTime),
			UpdatedAt:  ancestor.UpdatedAt.Format(time.DateTime),
		})
	}

	return &protobuf.GetCategoryAncestorsResponse{
		Ancestors: protoAncestors,
	}, nil
}

func (s Service) GetCategoryDescendants(ctx context.Context, req *protobuf.GetCategoryDescendantsRequest) (*protobuf.GetCategoryDescendantsResponse, error) {
	descendants, err := s.repository.GetDescendants(ctx, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get category descendants: %v", err)
	}

	var protoDescendants []*protobuf.GetCategoryDescendantsItem
	for _, descendant := range descendants {
		protoDescendants = append(protoDescendants, &protobuf.GetCategoryDescendantsItem{
			CategoryId: descendant.ID,
			Title:      descendant.Title,
			Slug:       descendant.Slug,
			CreatedAt:  descendant.CreatedAt.Format(time.DateTime),
			UpdatedAt:  descendant.UpdatedAt.Format(time.DateTime),
		})
	}

	return &protobuf.GetCategoryDescendantsResponse{
		Descendants: protoDescendants,
	}, nil
}

func (s Service) AddCategoryHierarchy(ctx context.Context, req *protobuf.AddCategoryHierarchyRequest) (*protobuf.AddCategoryHierarchyResponse, error) {
	err := s.repository.AddHierarchy(ctx, req.AncestorId, req.DescendantId, int(req.Depth))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add category hierarchy: %v", err)
	}

	return &protobuf.AddCategoryHierarchyResponse{
		Success: true,
	}, nil
}
