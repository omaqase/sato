package product

import "time"

type Entity struct {
	ID         string
	Title      string
	CategoryID string
	SellerID   string
	Stock      int
	Rating     float64
	Approved   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
