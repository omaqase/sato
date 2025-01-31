package domain

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       float64
	CategoryID  uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
