package category

import "time"

type Entity struct {
	ID        string
	Title     string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
