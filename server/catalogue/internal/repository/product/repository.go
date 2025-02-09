package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/catalogue/internal/domain/product"
)

type IRepository interface {
	Create(ctx context.Context, contract *CreateProductContract) (*product.Entity, error)
	Update(ctx context.Context, contract *UpdateProductContract) (*product.Entity, error)
	GetByID(ctx context.Context, id string) (*product.Entity, error)
	DeleteByID(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*product.Entity, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) IRepository {
	return &Repository{pool: pool}
}

func (r Repository) Create(ctx context.Context, contract *CreateProductContract) (*product.Entity, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				err = fmt.Errorf("transaction rollback error: %v, original error: %w", rbErr, err)
			}
			return
		}
		err = tx.Commit(ctx)
		if err != nil {
			err = fmt.Errorf("failed to commit transaction: %w", err)
		}
	}()

	dest := new(product.Entity)
	row := tx.QueryRow(ctx, CreateProductSQL, contract.Title, contract.CategoryID, contract.SellerID, contract.Stock)
	err = row.Scan(&dest.ID, &dest.Title, &dest.CategoryID, &dest.SellerID, &dest.Stock, &dest.Rating, &dest.Approved, &dest.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}
	return dest, nil
}

func (r Repository) Update(ctx context.Context, contract *UpdateProductContract) (*product.Entity, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				err = fmt.Errorf("transaction rollback error: %v, original error: %w", rbErr, err)
			}
			return
		}
		err = tx.Commit(ctx)
		if err != nil {
			err = fmt.Errorf("failed to commit transaction: %w", err)
		}
	}()

	dest := new(product.Entity)
	row := tx.QueryRow(ctx, UpdateProductSQL, contract.Title, contract.CategoryID, contract.SellerID, contract.Stock, contract.ID)
	err = row.Scan(&dest.ID, &dest.Title, &dest.CategoryID, &dest.SellerID, &dest.Stock, &dest.Rating, &dest.Approved, &dest.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	return dest, nil
}

func (r Repository) GetByID(ctx context.Context, id string) (*product.Entity, error) {
	dest := new(product.Entity)
	row := r.pool.QueryRow(ctx, GetProductByIDSQL, id)
	err := row.Scan(&dest.ID, &dest.Title, &dest.CategoryID, &dest.SellerID, &dest.Stock, &dest.Rating, &dest.Approved, &dest.CreatedAt, &dest.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("product with id %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to get product by id: %w", err)
	}
	return dest, nil
}

func (r Repository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, DeleteProductByIDSQL, id)
	if err != nil {
		return fmt.Errorf("failed to delete product by id: %w", err)
	}
	return nil
}

func (r Repository) List(ctx context.Context, limit, offset int) ([]*product.Entity, error) {
	rows, err := r.pool.Query(ctx, ListProductsSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	defer rows.Close()

	var products []*product.Entity
	for rows.Next() {
		dest := new(product.Entity)
		err := rows.Scan(&dest.ID, &dest.Title, &dest.CategoryID, &dest.SellerID, &dest.Stock, &dest.Rating, &dest.Approved, &dest.CreatedAt, &dest.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return products, nil
}
