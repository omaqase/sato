package domain

import "time"

type User struct {
	ID              string
	Email           string
	Role            string
	FirstName       string
	LastName        string
	Phone           string
	PaymentMethodID string
	Promotions      bool
	EmailVerified   bool
	TwoFactor       bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
