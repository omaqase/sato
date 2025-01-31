package product

import "github.com/google/uuid"

type CreateProductContract struct {
	Title       string
	Description string
	Price       float64
	CategoryID  uuid.UUID
}

func NewCreateProductContract(title, description string, price float64, categoryID uuid.UUID) *CreateProductContract {
	return &CreateProductContract{
		Title:       title,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}
}

type UpdateProductContract struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       float64
	CategoryID  uuid.UUID
}

func NewUpdateProductContract(id uuid.UUID, title, description string, price float64, categoryID uuid.UUID) *UpdateProductContract {
	return &UpdateProductContract{
		ID:          id,
		Title:       title,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}
}

type DeleteProductContract struct {
	ID uuid.UUID
}

func NewDeleteProductContract(id uuid.UUID) *DeleteProductContract {
	return &DeleteProductContract{ID: id}
}

type ListProductsContract struct {
	Page  int
	Limit int
}

func NewListProductsContract(page, limit int) *ListProductsContract {
	return &ListProductsContract{Page: page, Limit: limit}
}
