package service

import (
	"context"
	"errors"
	"fmt"

	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
)

// CustomerService handles customer business logic.
type CustomerService interface {
	List(ctx context.Context) ([]domain.Customer, error)
	GetByID(ctx context.Context, id int) (*domain.Customer, error)
	Create(ctx context.Context, c *domain.Customer) error
	Update(ctx context.Context, c *domain.Customer) error
}

type customerService struct {
	repo repository.CustomerRepository
}

// NewCustomerService creates a CustomerService.
func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) List(ctx context.Context) ([]domain.Customer, error) {
	return s.repo.List(ctx)
}

func (s *customerService) GetByID(ctx context.Context, id int) (*domain.Customer, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("customer tidak ditemukan")
	}
	return c, nil
}

func (s *customerService) Create(ctx context.Context, c *domain.Customer) error {
	if c.Nama == "" || c.Email == "" {
		return errors.New("nama dan email harus diisi")
	}
	if c.KodeCust == "" {
		return errors.New("kode_cust harus diisi")
	}

	emailExists, err := s.repo.EmailExists(ctx, c.Email, nil)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email sudah terdaftar")
	}

	kodeExists, err := s.repo.KodeCustExists(ctx, c.KodeCust, nil)
	if err != nil {
		return err
	}
	if kodeExists {
		return fmt.Errorf("kode_cust '%s' sudah terdaftar", c.KodeCust)
	}

	c.Aktif = true
	return s.repo.Create(ctx, c)
}

func (s *customerService) Update(ctx context.Context, c *domain.Customer) error {
	if c.Nama == "" || c.Email == "" {
		return errors.New("nama dan email harus diisi")
	}

	// Ensure customer exists
	if _, err := s.repo.FindByID(ctx, c.ID); err != nil {
		return errors.New("customer tidak ditemukan")
	}

	emailExists, err := s.repo.EmailExists(ctx, c.Email, &c.ID)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email sudah digunakan customer lain")
	}

	kodeExists, err := s.repo.KodeCustExists(ctx, c.KodeCust, &c.ID)
	if err != nil {
		return err
	}
	if kodeExists {
		return fmt.Errorf("kode_cust '%s' sudah digunakan customer lain", c.KodeCust)
	}

	return s.repo.Update(ctx, c)
}
