package category

import "github.com/google/uuid"

type CreateCategoryContract struct {
	Title string
	Slug  string
}

func NewCreateCategoryContract(title, slug string) *CreateCategoryContract {
	return &CreateCategoryContract{Title: title, Slug: slug}
}

type UpdateCategoryContract struct {
	ID    uuid.UUID
	Title string
	Slug  string
}

func NewUpdateCategoryContract(id uuid.UUID, title, slug string) *UpdateCategoryContract {
	return &UpdateCategoryContract{ID: id, Title: title, Slug: slug}
}

type DeleteCategoryContract struct {
	ID   *uuid.UUID
	Slug *string
}

func NewDeleteCategoryContractByID(id uuid.UUID) *DeleteCategoryContract {
	return &DeleteCategoryContract{ID: &id}
}

func NewDeleteCategoryContractBySlug(slug string) *DeleteCategoryContract {
	return &DeleteCategoryContract{Slug: &slug}
}

type ListCategoriesContract struct {
	Page  int
	Limit int
}

func NewListCategoriesContract(page, limit int) *ListCategoriesContract {
	return &ListCategoriesContract{Page: page, Limit: limit}
}
