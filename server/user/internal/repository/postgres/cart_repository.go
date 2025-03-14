package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/user/internal/domain"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) domain.ICartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddToCart(ctx context.Context, item domain.CartItem) error {
	query := `
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id) 
		DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity
	`
	_, err := r.db.Exec(ctx, query, item.UserID, item.ProductID, item.Quantity)
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}
	return nil
}

func (r *CartRepository) GetCart(ctx context.Context, userID uuid.UUID) ([]domain.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity, created_at, updated_at
		FROM cart_items
		WHERE user_id = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}
	defer rows.Close()

	var items []domain.CartItem
	for rows.Next() {
		var item domain.CartItem
		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.ProductID,
			&item.Quantity,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *CartRepository) UpdateCartItem(ctx context.Context, userID, productID uuid.UUID, quantity int) error {
	query := `
		UPDATE cart_items
		SET quantity = $3, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $1 AND product_id = $2
	`
	result, err := r.db.Exec(ctx, query, userID, productID, quantity)
	if err != nil {
		return fmt.Errorf("failed to update cart item: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("cart item not found")
	}
	return nil
}

func (r *CartRepository) RemoveFromCart(ctx context.Context, userID, productID uuid.UUID) error {
	query := `
		DELETE FROM cart_items
		WHERE user_id = $1 AND product_id = $2
	`
	result, err := r.db.Exec(ctx, query, userID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove item from cart: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("cart item not found")
	}
	return nil
}

func (r *CartRepository) ClearCart(ctx context.Context, userID uuid.UUID) error {
	query := `
		DELETE FROM cart_items
		WHERE user_id = $1
	`
	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to clear cart: %w", err)
	}
	return nil
}
