package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"insightflow/be-penjualan/internal/database"
	"insightflow/be-penjualan/internal/domain"
)

// ProdukRepository defines database operations for produk.
type ProdukRepository interface {
	List(ctx context.Context, aktifOnly bool) ([]domain.Produk, error)
	FindByID(ctx context.Context, id int) (*domain.Produk, error)
	Create(ctx context.Context, p *domain.Produk) error
	Update(ctx context.Context, p *domain.Produk) error
	Deactivate(ctx context.Context, id int) error
	KodeProdukExists(ctx context.Context, kode string, excludeID *int) (bool, error)
}

type produkRepo struct{}

// NewProdukRepository creates a new ProdukRepository backed by pgxpool.
func NewProdukRepository() ProdukRepository {
	return &produkRepo{}
}

func (r *produkRepo) List(ctx context.Context, aktifOnly bool) ([]domain.Produk, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	q := `
		SELECT id, kode_produk, nama, kategori_pakaian, ukuran, warna, bahan, harga, stok, aktif, created_at, updated_at
		FROM bisnis.tbl_produk
	`
	if aktifOnly {
		q += ` WHERE aktif = true`
	}
	q += ` ORDER BY nama ASC`

	rows, err := database.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Produk
	for rows.Next() {
		var p domain.Produk
		if err := rows.Scan(
			&p.ID, &p.KodeProduk, &p.Nama, &p.KategoriPakaian,
			&p.Ukuran, &p.Warna, &p.Bahan, &p.Harga, &p.Stok,
			&p.Aktif, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func (r *produkRepo) FindByID(ctx context.Context, id int) (*domain.Produk, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, kode_produk, nama, kategori_pakaian, ukuran, warna, bahan, harga, stok, aktif, created_at, updated_at
		FROM bisnis.tbl_produk
		WHERE id = $1
	`
	var p domain.Produk
	err := database.Pool.QueryRow(ctx, q, id).Scan(
		&p.ID, &p.KodeProduk, &p.Nama, &p.KategoriPakaian,
		&p.Ukuran, &p.Warna, &p.Bahan, &p.Harga, &p.Stok,
		&p.Aktif, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("produk not found")
		}
		return nil, err
	}
	return &p, nil
}

func (r *produkRepo) Create(ctx context.Context, p *domain.Produk) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		INSERT INTO bisnis.tbl_produk (kode_produk, nama, kategori_pakaian, ukuran, warna, bahan, harga, stok, aktif)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`
	return database.Pool.QueryRow(ctx, q,
		p.KodeProduk, p.Nama, p.KategoriPakaian, p.Ukuran, p.Warna,
		p.Bahan, p.Harga, p.Stok, p.Aktif,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *produkRepo) Update(ctx context.Context, p *domain.Produk) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		UPDATE bisnis.tbl_produk
		SET kode_produk = $1, nama = $2, kategori_pakaian = $3, ukuran = $4,
		    warna = $5, bahan = $6, harga = $7, stok = $8, aktif = $9,
		    updated_at = now()
		WHERE id = $10
		RETURNING updated_at
	`
	err := database.Pool.QueryRow(ctx, q,
		p.KodeProduk, p.Nama, p.KategoriPakaian, p.Ukuran,
		p.Warna, p.Bahan, p.Harga, p.Stok, p.Aktif, p.ID,
	).Scan(&p.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("produk not found")
		}
		return err
	}
	return nil
}

func (r *produkRepo) Deactivate(ctx context.Context, id int) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `UPDATE bisnis.tbl_produk SET aktif = false, updated_at = now() WHERE id = $1`
	tag, err := database.Pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("produk not found")
	}
	return nil
}

func (r *produkRepo) KodeProdukExists(ctx context.Context, kode string, excludeID *int) (bool, error) {
	if database.Pool == nil {
		return false, errors.New("database pool not initialised")
	}

	var exists bool
	var err error
	if excludeID == nil {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM bisnis.tbl_produk WHERE kode_produk = $1)`, kode,
		).Scan(&exists)
	} else {
		err = database.Pool.QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM bisnis.tbl_produk WHERE kode_produk = $1 AND id != $2)`, kode, *excludeID,
		).Scan(&exists)
	}
	return exists, err
}
