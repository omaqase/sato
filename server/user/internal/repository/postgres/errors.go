package postgres

import "errors"

var (
	ErrDatabaseInternalError    = errors.New("internal database error")
	ErrInvalidSignUpCredentials = errors.New("invalid sign up credentials")
	ErrInvalidCredentials       = errors.New("invalid credentials")
	ErrUserNotFound             = errors.New("user not found")
	ErrValidationFailed         = errors.New("validation failed")
	ErrAlreadyDeleted           = errors.New("user already deleted")
)
