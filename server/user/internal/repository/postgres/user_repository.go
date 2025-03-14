package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/user/internal/domain"
)

type IUsersRepository interface {
	Create(ctx context.Context, contract *CreateContract) (*domain.User, error)
	Update(ctx context.Context, contract *UpdateContract) (*domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, contract *GetByEmailContract) (*domain.User, error)
	GetSubscribedToPromotion(ctx context.Context, limit, offset int) ([]*domain.User, error)
	MakeSeller(ctx context.Context, id string) (*domain.User, error)
	AddToFavorites(ctx context.Context, contract *AddToFavoritesContract) (*domain.Favorite, error)
	RemoveFromFavorites(ctx context.Context, userID, productID string) error
	GetFavorites(ctx context.Context, contract *GetFavoritesContract) ([]*domain.Favorite, int, error)
}

type UsersRepository struct {
	pgx *pgxpool.Pool
}

func (r *UsersRepository) MakeSeller(ctx context.Context, id string) (*domain.User, error) {
	// Начинаем транзакцию
	tx, err := r.pgx.Begin(ctx)
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

	// Обновляем роль пользователя на "seller"
	updateRoleSQL := `
        UPDATE user_management.users
        SET role = 'seller'
        WHERE id = $1 AND deleted_at IS NULL
        RETURNING id, email, role, first_name, last_name, phone, promotions
    `
	user := new(domain.User)
	err = tx.QueryRow(ctx, updateRoleSQL, id).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Promotions,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user with id %s not found or deleted", id)
		}
		return nil, fmt.Errorf("failed to update user role: %w", err)
	}

	// Создаем запись в таблице sellers
	createSellerSQL := `
        INSERT INTO user_management.sellers(user_id, name, rating)
        VALUES ($1, 'test', $2)
    `
	_, err = tx.Exec(ctx, createSellerSQL, id, 0) // Начальный рейтинг = 0
	if err != nil {
		return nil, fmt.Errorf("failed to create seller record: %w", err)
	}

	return user, nil
}

func (r *UsersRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	user := new(domain.User)

	err := r.pgx.QueryRow(ctx, getUserByIDSQL, id).Scan(&user.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

func NewUsersRepository(pgx *pgxpool.Pool) IUsersRepository {
	return &UsersRepository{pgx: pgx}
}

func (r *UsersRepository) Create(ctx context.Context, contract *CreateContract) (*domain.User, error) {
	user := new(domain.User)

	err := r.pgx.QueryRow(ctx, `
		INSERT INTO user_management.users (email, role)
		VALUES ($1, 'customer')
		RETURNING id, email, role, first_name, last_name, phone, promotions
	`, contract.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Promotions,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
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

	var firstName, lastName, phone sql.NullString

	err := r.pgx.QueryRow(ctx, `
		SELECT 
			id, 
			email, 
			first_name, 
			last_name, 
			phone, 
			role, 
			promotions
		FROM user_management.users 
		WHERE email = $1 AND deleted_at IS NULL
	`, contract.Email).Scan(
		&user.ID,
		&user.Email,
		&firstName,
		&lastName,
		&phone,
		&user.Role,
		&user.Promotions,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	user.FirstName = firstName.String
	user.LastName = lastName.String
	user.Phone = &phone.String

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

func (r *UsersRepository) AddToFavorites(ctx context.Context, contract *AddToFavoritesContract) (*domain.Favorite, error) {
	favorite := new(domain.Favorite)

	err := r.pgx.QueryRow(ctx, addToFavoritesSQL, contract.UserID, contract.ProductID).Scan(
		&favorite.ID,
		&favorite.UserID,
		&favorite.ProductID,
		&favorite.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add to favorites: %w", err)
	}

	return favorite, nil
}

func (r *UsersRepository) RemoveFromFavorites(ctx context.Context, userID, productID string) error {
	result, err := r.pgx.Exec(ctx, removeFromFavoritesSQL, userID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove from favorites: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("favorite not found")
	}

	return nil
}

func (r *UsersRepository) GetFavorites(ctx context.Context, contract *GetFavoritesContract) ([]*domain.Favorite, int, error) {
	var total int
	err := r.pgx.QueryRow(ctx, getFavoritesCountSQL, contract.UserID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get favorites count: %w", err)
	}

	rows, err := r.pgx.Query(ctx, getFavoritesSQL, contract.UserID, contract.Limit, contract.Offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get favorites: %w", err)
	}
	defer rows.Close()

	var favorites []*domain.Favorite
	for rows.Next() {
		favorite := new(domain.Favorite)
		err := rows.Scan(
			&favorite.ID,
			&favorite.UserID,
			&favorite.ProductID,
			&favorite.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan favorite: %w", err)
		}
		favorites = append(favorites, favorite)
	}

	return favorites, total, nil
}
