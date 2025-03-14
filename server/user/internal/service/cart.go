package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/omaqase/sato/user/internal/domain"
)

type cartService struct {
	repo domain.ICartRepository
}

func NewCartService(repo domain.ICartRepository) domain.ICartService {
	return &cartService{repo: repo}
}

func (s *cartService) AddToCart(ctx context.Context, item domain.CartItem) error {
	return s.repo.AddToCart(ctx, item)
}

func (s *cartService) GetCart(ctx context.Context, userID uuid.UUID) ([]domain.CartItem, error) {
	return s.repo.GetCart(ctx, userID)
}

func (s *cartService) UpdateCartItem(ctx context.Context, userID, productID uuid.UUID, quantity int) error {
	return s.repo.UpdateCartItem(ctx, userID, productID, quantity)
}

func (s *cartService) RemoveFromCart(ctx context.Context, userID, productID uuid.UUID) error {
	return s.repo.RemoveFromCart(ctx, userID, productID)
}

func (s *cartService) ClearCart(ctx context.Context, userID uuid.UUID) error {
	return s.repo.ClearCart(ctx, userID)
}
