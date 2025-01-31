package domain

import "time"

type User struct {
	ID                       string
	Email                    string
	Password                 string
	FirstName                string
	LastName                 string
	Phone                    string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                time.Time
	IsSubscribedToPromotions bool
}

func NewUser(ID string, email string, password string, firstName string, lastName string, phone string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) *User {
	return &User{ID: ID, Email: email, Password: password, FirstName: firstName, LastName: lastName, Phone: phone, CreatedAt: createdAt, UpdatedAt: updatedAt, DeletedAt: deletedAt}
}
