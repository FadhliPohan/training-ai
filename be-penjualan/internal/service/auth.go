package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
)

// AuthService handles authentication business logic.
type AuthService interface {
	Login(ctx context.Context, email, password string) (*domain.User, string, time.Time, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Register(ctx context.Context, nama, email, password, role string, telegramID *int64) (*domain.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates an AuthService with the given user repository.
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(ctx context.Context, email, password string) (*domain.User, string, time.Time, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", time.Time{}, errors.New("email atau password salah")
	}
	if !user.Aktif {
		return nil, "", time.Time{}, errors.New("akun tidak aktif")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", time.Time{}, errors.New("email atau password salah")
	}

	token, expiresAt, err := generateJWT(user)
	if err != nil {
		return nil, "", time.Time{}, fmt.Errorf("gagal membuat token: %w", err)
	}
	return user, token, expiresAt, nil
}

func (s *authService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *authService) Register(ctx context.Context, nama, email, password, role string, telegramID *int64) (*domain.User, error) {
	exists, err := s.userRepo.EmailExists(ctx, email, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email sudah terdaftar")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("gagal hash password: %w", err)
	}

	user := &domain.User{
		Nama:           nama,
		Email:          email,
		Password:       string(hash),
		Role:           role,
		TelegramUserID: telegramID,
		Aktif:          true,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// generateJWT creates a signed JWT token for the given user (8 hour expiry).
func generateJWT(user *domain.User) (string, time.Time, error) {
	expiresAt := time.Now().Add(8 * time.Hour)
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
		"exp":     expiresAt.Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.App.JWTSecret))
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expiresAt, nil
}
