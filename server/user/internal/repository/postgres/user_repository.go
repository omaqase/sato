package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/user/internal/domain"
	"log"
	"time"
)

type IUsersRepository interface {
	Create(ctx context.Context, contract *CreateContract) (*domain.User, error)
	Update(ctx context.Context, contract *UpdateContract) (*domain.User, error)
	GetByEmail(ctx context.Context, contract *GetByEmailContract) (*domain.User, error)
	GetSubscribedToPromotion(ctx context.Context, limit, offset int) ([]*domain.User, error)
}

type UsersRepository struct {
	pgx *pgxpool.Pool
}

func NewUsersRepository(pgx *pgxpool.Pool) IUsersRepository {
	return &UsersRepository{pgx: pgx}
}

func (r *UsersRepository) Create(ctx context.Context, contract *CreateContract) (*domain.User, error) {
	tx, err := r.pgx.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
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

	user := new(domain.User)

	err = tx.QueryRow(ctx, createUserSQL, contract.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	return user, nil
}

func (r *UsersRepository) Update(ctx context.Context, contract *UpdateContract) (*domain.User, error) {
	tx, err := r.pgx.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
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

	now := time.Now().UTC()

	row := tx.QueryRow(ctx, updateUserSQL, contract.FirstName, contract.LastName, contract.Phone, contract.Promotions, now, contract.ID)

	user := new(domain.User)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Phone, &user.Promotions, &user.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user with id %s not found or deleted: %v", contract.ID, err)
		}
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return user, nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, contract *GetByEmailContract) (*domain.User, error) {
	user := new(domain.User)

	err := r.pgx.QueryRow(ctx, getUserByEmailSQL, contract.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return user, nil
}

func (r *UsersRepository) GetSubscribedToPromotion(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	rows, err := r.pgx.Query(ctx, getSubscribedToPromotionUsersSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := new(domain.User)
		if err := rows.Scan(&user.Email, &user.FirstName, &user.LastName); err != nil {
			log.Println(user)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rows: %v", err)
	}

	return users, nil
}
