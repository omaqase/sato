package domain

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID        uuid.UUID
	Title     string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
