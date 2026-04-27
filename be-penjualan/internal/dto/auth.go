package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== AUTH DTOs ====================

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"admin@insightflow.id"`
	Password string `json:"password" validate:"required,min=6" example:"Admin@12345"`
}

// LoginResponse represents the response payload for successful login
type LoginResponse struct {
	Token     string    `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresAt time.Time `json:"expires_at" example:"2026-04-25T01:18:10Z"`
	User      UserInfo  `json:"user"`
}

// LogoutRequest represents the request payload for user logout
type LogoutRequest struct {
	Token string `json:"token" validate:"required"`
}

// LogoutResponse represents the response payload for logout
type LogoutResponse struct {
	Message string `json:"message" example:"Logout berhasil"`
}

// RegisterRequest represents the request payload for user registration
type RegisterRequest struct {
	Nama     string `json:"nama" validate:"required,min=2,max=100" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=6" example:"SecurePass123"`
}

// RegisterResponse represents the response payload for successful registration
type RegisterResponse struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Nama      string    `json:"nama" example:"John Doe"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2026-04-24T10:00:00Z"`
}

// UserInfo contains non-sensitive user information
type UserInfo struct {
	ID    uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Nama  string    `json:"nama" example:"Administrator"`
	Email string    `json:"email" example:"admin@insightflow.id"`
	Role  string    `json:"role" example:"admin"`
}

// ProfileResponse represents the response payload for user profile
type ProfileResponse struct {
	ID             uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Nama           string     `json:"nama" example:"Administrator"`
	Email          string     `json:"email" example:"admin@insightflow.id"`
	Role           string     `json:"role" example:"admin"`
	TelegramUserID *int64     `json:"telegram_user_id,omitempty" example:"123456789"`
	Aktif          bool       `json:"aktif" example:"true"`
	CreatedAt      time.Time  `json:"created_at" example:"2026-04-24T10:00:00Z"`
}