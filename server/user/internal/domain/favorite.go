package domain

import (
	"time"
)

type Favorite struct {
	ID        string
	UserID    string
	ProductID string
	CreatedAt time.Time
}
