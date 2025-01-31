package postgres

type CreateUserContract struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func NewCreateUserContract(email, password, firstName, lastName string) *CreateUserContract {
	return &CreateUserContract{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
}

type UpdateUserContract struct {
	ID                       string
	FirstName                string
	LastName                 string
	Phone                    *string
	IsSubscribedToPromotions bool
}

func NewUpdateUserContract(
	id, firstName, lastName string,
	phone *string,
	isSubscribed bool,
) *UpdateUserContract {
	return &UpdateUserContract{
		ID:                       id,
		FirstName:                firstName,
		LastName:                 lastName,
		Phone:                    phone,
		IsSubscribedToPromotions: isSubscribed,
	}
}

type DeleteUserContract struct {
	ID string
}

func NewDeleteUserContract(id string) *DeleteUserContract {
	return &DeleteUserContract{ID: id}
}

type GetUserByCredentialsContract struct {
	Email    string
	Password string
}

func NewGetUserByCredentialsContract(email, password string) *GetUserByCredentialsContract {
	return &GetUserByCredentialsContract{Email: email, Password: password}
}

type GetSubscribedToPromotionContract struct {
	Limit  int
	Cursor string
}

func NewGetSubscribedToPromotionContract(limit int, cursor string) *GetSubscribedToPromotionContract {
	return &GetSubscribedToPromotionContract{Limit: limit, Cursor: cursor}
}
