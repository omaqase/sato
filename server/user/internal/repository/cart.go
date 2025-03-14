package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/user/internal/domain"
)

type cartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) domain.CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetCart(userID uuid.UUID) ([]domain.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity, created_at, updated_at
		FROM cart_items
		WHERE user_id = $1
	`

	rows, err := r.db.Query(context.Background(), query, userID)
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

func (r *cartRepository) AddToCart(userID uuid.UUID, productID uuid.UUID, quantity int32) (*domain.CartItem, error) {
	query := `
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id) DO UPDATE
		SET quantity = cart_items.quantity + EXCLUDED.quantity,
			updated_at = CURRENT_TIMESTAMP
		RETURNING id, user_id, product_id, quantity, created_at, updated_at
	`

	var item domain.CartItem
	err := r.db.QueryRow(
		context.Background(),
		query,
		userID,
		productID,
		quantity,
	).Scan(
		&item.ID,
		&item.UserID,
		&item.ProductID,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add to cart: %w", err)
	}

	return &item, nil
}

func (r *cartRepository) UpdateCartItem(userID uuid.UUID, productID uuid.UUID, quantity int32) (*domain.CartItem, error) {
	query := `
		UPDATE cart_items
		SET quantity = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $1 AND product_id = $2
		RETURNING id, user_id, product_id, quantity, created_at, updated_at
	`

	var item domain.CartItem
	err := r.db.QueryRow(
		context.Background(),
		query,
		userID,
		productID,
		quantity,
	).Scan(
		&item.ID,
		&item.UserID,
		&item.ProductID,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update cart item: %w", err)
	}

	return &item, nil
}

func (r *cartRepository) RemoveFromCart(userID uuid.UUID, productID uuid.UUID) error {
	query := `
		DELETE FROM cart_items
		WHERE user_id = $1 AND product_id = $2
	`

	result, err := r.db.Exec(context.Background(), query, userID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove from cart: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("cart item not found")
	}

	return nil
}

func (r *cartRepository) ClearCart(userID uuid.UUID) error {
	query := `
		DELETE FROM cart_items
		WHERE user_id = $1
	`

	_, err := r.db.Exec(context.Background(), query, userID)
	if err != nil {
		return fmt.Errorf("failed to clear cart: %w", err)
	}

	return nil
}
