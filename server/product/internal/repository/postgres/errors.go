package postgres

import "errors"

var (
	ErrDatabaseInternalError = errors.New("internal database error")
	ErrInvalidSlug           = errors.New("invalid or duplicate slug")
	ErrCategoryNotFound      = errors.New("category not found")
	ErrProductNotFound       = errors.New("product not found")
	ErrAlreadyDeleted        = errors.New("already deleted")
)
