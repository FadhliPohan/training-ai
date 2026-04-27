package service

import (
	"context"
	"errors"
	"fmt"

	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
)

// ProdukService handles product business logic.
type ProdukService interface {
	List(ctx context.Context, aktifOnly bool) ([]domain.Produk, error)
	GetByID(ctx context.Context, id int) (*domain.Produk, error)
	Create(ctx context.Context, p *domain.Produk) error
	Update(ctx context.Context, p *domain.Produk) error
	Deactivate(ctx context.Context, id int) error
}

type produkService struct {
	repo repository.ProdukRepository
}

// NewProdukService creates a ProdukService.
func NewProdukService(repo repository.ProdukRepository) ProdukService {
	return &produkService{repo: repo}
}

func (s *produkService) List(ctx context.Context, aktifOnly bool) ([]domain.Produk, error) {
	return s.repo.List(ctx, aktifOnly)
}

func (s *produkService) GetByID(ctx context.Context, id int) (*domain.Produk, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("produk tidak ditemukan")
	}
	return p, nil
}

func (s *produkService) Create(ctx context.Context, p *domain.Produk) error {
	if p.KodeProduk == "" || p.Nama == "" {
		return errors.New("kode_produk dan nama harus diisi")
	}
	if p.Harga <= 0 {
		return errors.New("harga harus lebih dari 0")
	}
	if p.Stok < 0 {
		return errors.New("stok tidak boleh negatif")
	}

	exists, err := s.repo.KodeProdukExists(ctx, p.KodeProduk, nil)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("kode_produk '%s' sudah terdaftar", p.KodeProduk)
	}

	p.Aktif = true
	return s.repo.Create(ctx, p)
}

func (s *produkService) Update(ctx context.Context, p *domain.Produk) error {
	if p.KodeProduk == "" || p.Nama == "" {
		return errors.New("kode_produk dan nama harus diisi")
	}
	if p.Harga <= 0 {
		return errors.New("harga harus lebih dari 0")
	}
	if p.Stok < 0 {
		return errors.New("stok tidak boleh negatif")
	}

	// Ensure product exists first
	if _, err := s.repo.FindByID(ctx, p.ID); err != nil {
		return errors.New("produk tidak ditemukan")
	}

	exists, err := s.repo.KodeProdukExists(ctx, p.KodeProduk, &p.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("kode_produk '%s' sudah digunakan produk lain", p.KodeProduk)
	}

	return s.repo.Update(ctx, p)
}

func (s *produkService) Deactivate(ctx context.Context, id int) error {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		return errors.New("produk tidak ditemukan")
	}
	return s.repo.Deactivate(ctx, id)
}
