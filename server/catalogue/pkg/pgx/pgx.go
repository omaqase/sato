package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGXPool(ctx context.Context, config *pgxpool.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
