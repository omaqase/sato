package category

type CreateCategoryContract struct {
	Title string
	Slug  string
}

func NewCreateCategoryContract(title string, slug string) *CreateCategoryContract {
	return &CreateCategoryContract{Title: title, Slug: slug}
}

type UpdateCategoryContract struct {
	ID    string
	Title string
	Slug  string
}

func NewUpdateCategoryContract(ID string, title string, slug string) *UpdateCategoryContract {
	return &UpdateCategoryContract{ID: ID, Title: title, Slug: slug}
}
