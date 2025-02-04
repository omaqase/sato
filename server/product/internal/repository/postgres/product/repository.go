package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/product/internal/domain"
	"github.com/omaqase/sato/product/internal/repository/postgres"
	"time"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, contract *CreateProductContract) (*domain.Product, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	UpdateProduct(ctx context.Context, contract *UpdateProductContract) (*domain.Product, error)
	DeleteProduct(ctx context.Context, contract *DeleteProductContract) error
	ListProducts(ctx context.Context, contract *ListProductsContract) ([]*domain.Product, int, error)
}

type ProductRepository struct {
	pgx *pgxpool.Pool
}

func NewProductRepository(pgx *pgxpool.Pool) IProductRepository {
	return &ProductRepository{pgx: pgx}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, contract *CreateProductContract) (*domain.Product, error) {
	product := &domain.Product{}
	now := time.Now().UTC()
	err := r.pgx.QueryRow(ctx, createProductQuery,
		contract.Title,
		contract.Description,
		contract.Price,
		contract.CategoryID,
		now,
		now,
	).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	return product, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	product := &domain.Product{}
	err := r.pgx.QueryRow(ctx, getProductByIDQuery, id).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, postgres.ErrProductNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, contract *UpdateProductContract) (*domain.Product, error) {
	product := &domain.Product{}
	err := r.pgx.QueryRow(ctx, updateProductQuery,
		contract.Title,
		contract.Description,
		contract.Price,
		contract.CategoryID,
		contract.ID,
	).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, postgres.ErrProductNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	return product, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, contract *DeleteProductContract) error {
	result, err := r.pgx.Exec(ctx, deleteProductQuery, contract.ID)
	if err != nil {
		return fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	if result.RowsAffected() == 0 {
		return postgres.ErrAlreadyDeleted
	}
	return nil
}

func (r *ProductRepository) ListProducts(ctx context.Context, contract *ListProductsContract) ([]*domain.Product, int, error) {
	offset := (contract.Page - 1) * contract.Limit
	rows, err := r.pgx.Query(ctx, listProductsQuery, contract.Limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	var totalCount int
	err = r.pgx.QueryRow(ctx, `
        SELECT COUNT(*)
        FROM product_management.products
        WHERE deleted_at IS NULL`).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	return products, totalCount, nil
}
