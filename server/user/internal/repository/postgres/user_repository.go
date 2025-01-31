package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/user/internal/domain"
	"time"
)

type IUsersRepository interface {
	Create(ctx context.Context, contract *CreateUserContract) (*domain.User, error)
	Update(ctx context.Context, contract *UpdateUserContract) (*domain.User, error)
	Delete(ctx context.Context, contract *DeleteUserContract) error
	GetByCredentials(ctx context.Context, contract *GetUserByCredentialsContract) (*domain.User, error)
	GetSubscribedToPromotions(ctx context.Context, contract *GetSubscribedToPromotionContract) ([]string, error)
}

type UsersRepository struct {
	pgx *pgxpool.Pool
}

func NewUsersRepository(pgx *pgxpool.Pool) IUsersRepository {
	return &UsersRepository{pgx: pgx}
}

func (r *UsersRepository) Create(ctx context.Context, contract *CreateUserContract) (*domain.User, error) {
	user := &domain.User{}
	now := time.Now().UTC()

	err := r.pgx.QueryRow(ctx, createUserQuery,
		contract.Email,
		contract.Password,
		contract.FirstName,
		contract.LastName,
		now,
		now,
	).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsSubscribedToPromotions,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return nil, ErrInvalidSignUpCredentials
			case "23514":
				return nil, ErrValidationFailed
			}
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseInternalError, err)
	}

	return user, nil
}

func (r *UsersRepository) Update(ctx context.Context, contract *UpdateUserContract) (*domain.User, error) {
	user := &domain.User{}

	err := r.pgx.QueryRow(ctx, updateUserQuery,
		contract.FirstName,
		contract.LastName,
		contract.Phone,
		contract.IsSubscribedToPromotions,
		contract.ID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsSubscribedToPromotions,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return nil, ErrInvalidSignUpCredentials
			case "22P02":
				return nil, ErrValidationFailed
			}
		}

		return nil, fmt.Errorf("%w: %v", ErrDatabaseInternalError, err)
	}

	return user, nil
}

func (r *UsersRepository) Delete(ctx context.Context, contract *DeleteUserContract) error {
	var deletedAt time.Time

	err := r.pgx.QueryRow(ctx, deleteUserQuery, contract.ID).Scan(&deletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrAlreadyDeleted
		}

		return fmt.Errorf("%w: %v", ErrDatabaseInternalError, err)
	}

	return nil
}

func (r *UsersRepository) GetByCredentials(ctx context.Context, contract *GetUserByCredentialsContract) (*domain.User, error) {
	var user domain.User

	err := r.pgx.QueryRow(ctx, getUserByCredentialsQuery, contract.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}

		return nil, ErrDatabaseInternalError
	}

	if user.Password != contract.Password {
		return nil, ErrInvalidCredentials
	}

	return &user, nil
}

func (r *UsersRepository) GetSubscribedToPromotions(ctx context.Context, contract *GetSubscribedToPromotionContract) ([]string, error) {
	rows, err := r.pgx.Query(ctx, getUserSubscribedToPromotions, contract.Cursor, contract.Limit)
	if err != nil {
		return nil, ErrDatabaseInternalError
	}
	defer rows.Close()

	emails := make([]string, 0, contract.Limit)
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, ErrDatabaseInternalError
		}
		emails = append(emails, email)
	}

	if err = rows.Err(); err != nil {
		return nil, ErrDatabaseInternalError
	}

	return emails, nil
}
