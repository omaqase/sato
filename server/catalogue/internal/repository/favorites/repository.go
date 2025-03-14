// server/catalogue/internal/repository/favorites/repository.go
package favorites

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddToFavorites(ctx context.Context, userId, productId string) error {
	_, err := r.db.Exec(ctx, `
        INSERT INTO product_management.user_favorites (user_id, product_id)
        VALUES ($1, $2)
        ON CONFLICT (user_id, product_id) DO NOTHING
    `, userId, productId)
	return err
}

func (r *Repository) RemoveFromFavorites(ctx context.Context, userId, productId string) error {
	_, err := r.db.Exec(ctx, `
        DELETE FROM product_management.user_favorites
        WHERE user_id = $1 AND product_id = $2
    `, userId, productId)
	return err
}

func (r *Repository) GetUserFavorites(ctx context.Context, userId string, limit, offset int) ([]string, error) {
	rows, err := r.db.Query(ctx, `
        SELECT product_id 
        FROM product_management.user_favorites
        WHERE user_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productIds []string
	for rows.Next() {
		var productId string
		if err := rows.Scan(&productId); err != nil {
			return nil, err
		}
		productIds = append(productIds, productId)
	}
	return productIds, nil
}
