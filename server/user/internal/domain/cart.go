package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ICartRepository interface {
	AddToCart(ctx context.Context, item CartItem) error
	GetCart(ctx context.Context, userID uuid.UUID) ([]CartItem, error)
	UpdateCartItem(ctx context.Context, userID, productID uuid.UUID, quantity int) error
	RemoveFromCart(ctx context.Context, userID, productID uuid.UUID) error
	ClearCart(ctx context.Context, userID uuid.UUID) error
}

type ICartService interface {
	AddToCart(ctx context.Context, item CartItem) error
	GetCart(ctx context.Context, userID uuid.UUID) ([]CartItem, error)
	UpdateCartItem(ctx context.Context, userID, productID uuid.UUID, quantity int) error
	RemoveFromCart(ctx context.Context, userID, productID uuid.UUID) error
	ClearCart(ctx context.Context, userID uuid.UUID) error
}
