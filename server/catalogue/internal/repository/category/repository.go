package category

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/catalogue/internal/domain/category"
)

type IRepository interface {
	Create(ctx context.Context, contract *CreateCategoryContract) (*category.Entity, error)
	Update(ctx context.Context, contract *UpdateCategoryContract) (*category.Entity, error)
	GetByID(ctx context.Context, id string) (*category.Entity, error)
	GetBySlug(ctx context.Context, slug string) (*category.Entity, error)
	List(ctx context.Context, limit, offset int) ([]*category.Entity, error)
	DeleteByID(ctx context.Context, id string) error
	DeleteBySlug(ctx context.Context, slug string) error
	AddHierarchy(ctx context.Context, ancestorID, descendantID string, depth int) error
	GetAncestors(ctx context.Context, categoryID string) ([]*category.Entity, error)
	GetDescendants(ctx context.Context, categoryID string) ([]*category.Entity, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) IRepository {
	return &Repository{pool: pool}
}

func (r Repository) Create(ctx context.Context, contract *CreateCategoryContract) (*category.Entity, error) {
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

	dest := new(category.Entity)
	row := tx.QueryRow(ctx, CreateCategorySQL, contract.Title, contract.Slug)
	err = row.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}
	return dest, nil
}

func (r Repository) Update(ctx context.Context, contract *UpdateCategoryContract) (*category.Entity, error) {
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

	updatedAt := time.Now()
	row := tx.QueryRow(ctx, UpdateCategorySQL, contract.Title, contract.Slug, updatedAt, contract.ID)
	dest := new(category.Entity)
	err = row.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}
	return dest, nil
}

func (r Repository) GetByID(ctx context.Context, id string) (*category.Entity, error) {
	dest := new(category.Entity)
	row := r.pool.QueryRow(ctx, GetCategoryByIDSQL, id)
	err := row.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.CreatedAt, &dest.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("category with id %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}
	return dest, nil
}

func (r Repository) GetBySlug(ctx context.Context, slug string) (*category.Entity, error) {
	dest := new(category.Entity)
	row := r.pool.QueryRow(ctx, GetCategoryBySlugSQL, slug)
	err := row.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.CreatedAt, &dest.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("category with slug %s not found: %w", slug, err)
		}
		return nil, fmt.Errorf("failed to get category by slug: %w", err)
	}
	return dest, nil
}

func (r Repository) List(ctx context.Context, limit, offset int) ([]*category.Entity, error) {
	rows, err := r.pool.Query(ctx, ListCategoriesSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	defer rows.Close()

	var categories []*category.Entity
	for rows.Next() {
		dest := new(category.Entity)
		err := rows.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.CreatedAt, &dest.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return categories, nil
}

func (r Repository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, DeleteCategoryByIDSQL, id)
	if err != nil {
		return fmt.Errorf("failed to delete category by id: %w", err)
	}
	return nil
}

func (r Repository) DeleteBySlug(ctx context.Context, slug string) error {
	_, err := r.pool.Exec(ctx, DeleteCategoryBySlugSQL, slug)
	if err != nil {
		return fmt.Errorf("failed to delete category by slug: %w", err)
	}
	return nil
}

func (r Repository) AddHierarchy(ctx context.Context, ancestorID, descendantID string, depth int) error {
	_, err := r.pool.Exec(ctx, AddCategoryHierarchySQL, ancestorID, descendantID, depth)
	if err != nil {
		return fmt.Errorf("failed to add category hierarchy: %w", err)
	}
	return nil
}

func (r Repository) GetAncestors(ctx context.Context, categoryID string) ([]*category.Entity, error) {
	rows, err := r.pool.Query(ctx, GetCategoryAncestorsSQL, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get category ancestors: %w", err)
	}
	defer rows.Close()

	var categories []*category.Entity
	for rows.Next() {
		dest := new(category.Entity)
		err := rows.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.CreatedAt, &dest.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category ancestor: %w", err)
		}
		categories = append(categories, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return categories, nil
}

func (r Repository) GetDescendants(ctx context.Context, categoryID string) ([]*category.Entity, error) {
	rows, err := r.pool.Query(ctx, GetCategoryDescendantsSQL, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get category descendants: %w", err)
	}
	defer rows.Close()

	var categories []*category.Entity
	for rows.Next() {
		dest := new(category.Entity)
		err := rows.Scan(&dest.ID, &dest.Title, &dest.Slug, &dest.CreatedAt, &dest.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category descendant: %w", err)
		}
		categories = append(categories, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return categories, nil
}
