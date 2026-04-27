package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
)

// UserService handles user management business logic.
type UserService interface {
	List(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Create(ctx context.Context, u *domain.User, password string) error
	Update(ctx context.Context, u *domain.User) error
	Deactivate(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a UserService.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) List(ctx context.Context) ([]domain.User, error) {
	return s.repo.List(ctx)
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	return u, nil
}

func (s *userService) Create(ctx context.Context, u *domain.User, password string) error {
	if u.Nama == "" || u.Email == "" {
		return errors.New("nama dan email harus diisi")
	}
	if len(password) < 8 {
		return errors.New("password minimal 8 karakter")
	}
	if !isValidRole(u.Role) {
		return fmt.Errorf("role tidak valid: %s. Valid: admin, manager, sales, viewer", u.Role)
	}

	exists, err := s.repo.EmailExists(ctx, u.Email, nil)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email sudah terdaftar")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("gagal hash password: %w", err)
	}
	u.Password = string(hash)
	u.Aktif = true

	return s.repo.Create(ctx, u)
}

func (s *userService) Update(ctx context.Context, u *domain.User) error {
	if u.Nama == "" || u.Email == "" {
		return errors.New("nama dan email harus diisi")
	}
	if !isValidRole(u.Role) {
		return fmt.Errorf("role tidak valid: %s", u.Role)
	}

	// Ensure user exists
	if _, err := s.repo.FindByID(ctx, u.ID); err != nil {
		return errors.New("user tidak ditemukan")
	}

	exists, err := s.repo.EmailExists(ctx, u.Email, &u.ID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email sudah digunakan user lain")
	}

	return s.repo.Update(ctx, u)
}

func (s *userService) Deactivate(ctx context.Context, id uuid.UUID) error {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		return errors.New("user tidak ditemukan")
	}
	return s.repo.Deactivate(ctx, id)
}

func isValidRole(role string) bool {
	switch role {
	case domain.RoleAdmin, domain.RoleManager, domain.RoleSales, domain.RoleViewer:
		return true
	}
	return false
}
