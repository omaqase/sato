package product

type CreateProductContract struct {
	SellerID   string
	Title      string
	Slug       string
	CategoryID string
	Stock      int
}

func NewCreateProductContract(sellerID string, title string, slug string, categoryID string, stock int) *CreateProductContract {
	return &CreateProductContract{
		SellerID:   sellerID,
		Title:      title,
		Slug:       slug,
		CategoryID: categoryID,
		Stock:      stock,
	}
}

type UpdateProductContract struct {
	ID         string
	Title      string
	Slug       string
	CategoryID string
	SellerID   string
	Stock      int
}

func NewUpdateProductContract(id string, title string, slug string, categoryID string, stock int, sellerID string) *UpdateProductContract {
	return &UpdateProductContract{
		ID:         id,
		Title:      title,
		Slug:       slug,
		CategoryID: categoryID,
		SellerID:   sellerID,
		Stock:      stock,
	}
}
