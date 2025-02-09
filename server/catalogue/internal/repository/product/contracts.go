package product

type CreateProductContract struct {
	SellerID   string
	Title      string
	CategoryID string
	Stock      int
}

func NewCreateProductContract(sellerID string, title string, categoryID string, stock int) *CreateProductContract {
	return &CreateProductContract{
		SellerID:   sellerID,
		Title:      title,
		CategoryID: categoryID,
		Stock:      stock,
	}
}

type UpdateProductContract struct {
	ID         string
	Title      string
	CategoryID string
	SellerID   string
	Stock      int
}

func NewUpdateProductContract(id string, title string, categoryID string, sellerID string, stock int) *UpdateProductContract {
	return &UpdateProductContract{
		ID:         id,
		Title:      title,
		CategoryID: categoryID,
		SellerID:   sellerID,
		Stock:      stock,
	}
}
