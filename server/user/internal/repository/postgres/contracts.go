package postgres

type CreateContract struct {
	Email string
}

type GetByEmailContract struct {
	Email string
}

type UpdateContract struct {
	ID         string
	FirstName  string
	LastName   string
	Phone      string
	Promotions bool
}

type AddToFavoritesContract struct {
	UserID    string
	ProductID string
}

type GetFavoritesContract struct {
	UserID string
	Limit  int
	Offset int
}
