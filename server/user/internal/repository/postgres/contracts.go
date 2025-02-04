package postgres

type CreateContract struct {
	Email     string
	FirstName string
	LastName  string
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
