package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"insightflow/be-penjualan/internal/database"
	"insightflow/be-penjualan/internal/domain"
)

// CustomerRepository defines database operations for customers.
type CustomerRepository interface {
	List(ctx context.Context) ([]domain.Customer, error)
	FindByID(ctx context.Context, id int) (*domain.Customer, error)
	FindByEmail(ctx context.Context, email string) (*domain.Customer, error)
	Create(ctx context.Context, c *domain.Customer) error
	Update(ctx context.Context, c *domain.Customer) error
	EmailExists(ctx context.Context, email string, excludeID *int) (bool, error)
	KodeCustExists(ctx context.Context, kode string, excludeID *int) (bool, error)
}

type customerRepo struct{}

// NewCustomerRepository creates a new CustomerRepository backed by pgxpool.
func NewCustomerRepository() CustomerRepository {
	return &customerRepo{}
}

func (r *customerRepo) List(ctx context.Context) ([]domain.Customer, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, kode_cust, nama, email, telepon, alamat, aktif, created_at
		FROM bisnis.tbl_customer
		ORDER BY nama ASC
	`
	rows, err := database.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Customer
	for rows.Next() {
		var c domain.Customer
		if err := rows.Scan(
			&c.ID, &c.KodeCust, &c.Nama, &c.Email,
			&c.Telepon, &c.Alamat, &c.Aktif, &c.CreatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, rows.Err()
}

func (r *customerRepo) FindByID(ctx context.Context, id int) (*domain.Customer, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, kode_cust, nama, email, telepon, alamat, aktif, created_at
		FROM bisnis.tbl_customer
		WHERE id = $1
	`
	var c domain.Customer
	err := database.Pool.QueryRow(ctx, q, id).Scan(
		&c.ID, &c.KodeCust, &c.Nama, &c.Email,
		&c.Telepon, &c.Alamat, &c.Aktif, &c.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, err
	}
	return &c, nil
}

func (r *customerRepo) FindByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, kode_cust, nama, email, password, telepon, alamat, aktif, created_at
		FROM bisnis.tbl_customer
		WHERE email = $1
	`
	var c domain.Customer
	err := database.Pool.QueryRow(ctx, q, email).Scan(
		&c.ID, &c.KodeCust, &c.Nama, &c.Email, &c.Password,
		&c.Telepon, &c.Alamat, &c.Aktif, &c.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, err
	}
	return &c, nil
}

func (r *customerRepo) Create(ctx context.Context, c *domain.Customer) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		INSERT INTO bisnis.tbl_customer (kode_cust, nama, email, password, telepon, alamat, aktif)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`
	return database.Pool.QueryRow(ctx, q,
		c.KodeCust, c.Nama, c.Email, c.Password, c.Telepon, c.Alamat, c.Aktif,
	).Scan(&c.ID, &c.CreatedAt)
}

func (r *customerRepo) Update(ctx context.Context, c *domain.Customer) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		UPDATE bisnis.tbl_customer
		SET kode_cust = $1, nama = $2, email = $3, telepon = $4, alamat = $5, aktif = $6
		WHERE id = $7
	`
	tag, err := database.Pool.Exec(ctx, q,
		c.KodeCust, c.Nama, c.Email, c.Telepon, c.Alamat, c.Aktif, c.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("customer not found")
	}
	return nil
}

func (r *customerRepo) EmailExists(ctx context.Context, email string, excludeID *int) (bool, error) {
	if database.Pool == nil {
		return false, errors.New("database pool not initialised")
	}

	var exists bool
	var err error
	if excludeID == nil {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM bisnis.tbl_customer WHERE email = $1)`, email,
		).Scan(&exists)
	} else {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM bisnis.tbl_customer WHERE email = $1 AND id != $2)`, email, *excludeID,
		).Scan(&exists)
	}
	return exists, err
}

func (r *customerRepo) KodeCustExists(ctx context.Context, kode string, excludeID *int) (bool, error) {
	if database.Pool == nil {
		return false, errors.New("database pool not initialised")
	}

	var exists bool
	var err error
	if excludeID == nil {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM bisnis.tbl_customer WHERE kode_cust = $1)`, kode,
		).Scan(&exists)
	} else {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM bisnis.tbl_customer WHERE kode_cust = $1 AND id != $2)`, kode, *excludeID,
		).Scan(&exists)
	}
	return exists, err
}
