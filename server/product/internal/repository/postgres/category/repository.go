package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/product/internal/domain"
	"github.com/omaqase/sato/product/internal/repository/postgres"
	"time"
)

type ICategoryRepository interface {
	CreateCategory(ctx context.Context, contract *CreateCategoryContract) (*domain.Category, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*domain.Category, error)
	UpdateCategory(ctx context.Context, contract *UpdateCategoryContract) (*domain.Category, error)
	DeleteCategory(ctx context.Context, contract *DeleteCategoryContract) error
	ListCategories(ctx context.Context, contract *ListCategoriesContract) ([]*domain.Category, int, error)
}

type CategoryRepository struct {
	pgx *pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) ICategoryRepository {
	return &CategoryRepository{pgx: pgx}
}
func (r *CategoryRepository) CreateCategory(ctx context.Context, contract *CreateCategoryContract) (*domain.Category, error) {
	category := &domain.Category{}
	now := time.Now().UTC()

	err := r.pgx.QueryRow(ctx, createCategoryQuery,
		contract.Title,
		contract.Slug,
		now,
		now,
	).Scan(
		&category.ID,
		&category.Title,
		&category.Slug,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, postgres.ErrInvalidSlug
		}
		return nil, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	return category, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	category := &domain.Category{}

	err := r.pgx.QueryRow(ctx, getCategoryByIDQuery, id).Scan(
		&category.ID,
		&category.Title,
		&category.Slug,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, postgres.ErrCategoryNotFound
		}
		return nil, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, contract *UpdateCategoryContract) (*domain.Category, error) {
	category := &domain.Category{}

	err := r.pgx.QueryRow(ctx, updateCategoryQuery,
		contract.Title,
		contract.Slug,
		contract.ID,
	).Scan(
		&category.ID,
		&category.Title,
		&category.Slug,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, postgres.ErrCategoryNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, postgres.ErrInvalidSlug
		}
		return nil, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	return category, nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, contract *DeleteCategoryContract) error {
	var query string
	var args []interface{}

	if contract.ID != nil {
		query = deleteCategoryQuery
		args = []interface{}{*contract.ID}
	} else if contract.Slug != nil {
		query = `
            UPDATE product_management.categories
            SET deleted_at = NOW(), updated_at = NOW()
            WHERE slug = $1 AND deleted_at IS NULL`
		args = []interface{}{*contract.Slug}
	} else {
		return errors.New("either ID or Slug must be provided")
	}

	result, err := r.pgx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	if result.RowsAffected() == 0 {
		return postgres.ErrAlreadyDeleted
	}

	return nil
}

func (r *CategoryRepository) ListCategories(ctx context.Context, contract *ListCategoriesContract) ([]*domain.Category, int, error) {
	offset := (contract.Page - 1) * contract.Limit

	rows, err := r.pgx.Query(ctx, listCategoriesQuery, contract.Limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		category := &domain.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Title,
			&category.Slug,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	var totalCount int
	err = r.pgx.QueryRow(ctx, `
        SELECT COUNT(*)
        FROM product_management.categories
        WHERE deleted_at IS NULL`).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: %v", postgres.ErrDatabaseInternalError, err)
	}

	return categories, totalCount, nil
}
