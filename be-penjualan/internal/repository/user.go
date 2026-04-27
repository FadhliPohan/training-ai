package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"insightflow/be-penjualan/internal/database"
	"insightflow/be-penjualan/internal/domain"
)

// UserRepository defines database operations for users.
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Deactivate(ctx context.Context, id uuid.UUID) error
	EmailExists(ctx context.Context, email string, excludeID *uuid.UUID) (bool, error)
}

type userRepo struct{}

// NewUserRepository creates a new UserRepository backed by pgxpool.
func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, nama, email, password, role, telegram_user_id, aktif, created_at
		FROM app.users
		WHERE email = $1
	`
	var u domain.User
	err := database.Pool.QueryRow(ctx, q, email).Scan(
		&u.ID, &u.Nama, &u.Email, &u.Password, &u.Role,
		&u.TelegramUserID, &u.Aktif, &u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, nama, email, password, role, telegram_user_id, aktif, created_at
		FROM app.users
		WHERE id = $1
	`
	var u domain.User
	err := database.Pool.QueryRow(ctx, q, id).Scan(
		&u.ID, &u.Nama, &u.Email, &u.Password, &u.Role,
		&u.TelegramUserID, &u.Aktif, &u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) List(ctx context.Context) ([]domain.User, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, nama, email, role, telegram_user_id, aktif, created_at
		FROM app.users
		ORDER BY created_at DESC
	`
	rows, err := database.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(
			&u.ID, &u.Nama, &u.Email, &u.Role,
			&u.TelegramUserID, &u.Aktif, &u.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *userRepo) Create(ctx context.Context, u *domain.User) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		INSERT INTO app.users (nama, email, password, role, telegram_user_id, aktif)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	return database.Pool.QueryRow(ctx, q,
		u.Nama, u.Email, u.Password, u.Role, u.TelegramUserID, u.Aktif,
	).Scan(&u.ID, &u.CreatedAt)
}

func (r *userRepo) Update(ctx context.Context, u *domain.User) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		UPDATE app.users
		SET nama = $1, email = $2, role = $3, telegram_user_id = $4, aktif = $5
		WHERE id = $6
	`
	tag, err := database.Pool.Exec(ctx, q,
		u.Nama, u.Email, u.Role, u.TelegramUserID, u.Aktif, u.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *userRepo) Deactivate(ctx context.Context, id uuid.UUID) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `UPDATE app.users SET aktif = false WHERE id = $1`
	tag, err := database.Pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *userRepo) EmailExists(ctx context.Context, email string, excludeID *uuid.UUID) (bool, error) {
	if database.Pool == nil {
		return false, errors.New("database pool not initialised")
	}

	var exists bool
	var err error
	if excludeID == nil {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM app.users WHERE email = $1)`, email,
		).Scan(&exists)
	} else {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM app.users WHERE email = $1 AND id != $2)`, email, *excludeID,
		).Scan(&exists)
	}
	return exists, err
}
