package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== USER DTOs ====================

// UserRequest represents the request payload for creating/updating a user
type UserRequest struct {
	Nama           string  `json:"nama" validate:"required,min=2,max=100" example:"Jane Doe"`
	Email          string  `json:"email" validate:"required,email" example:"jane.doe@company.com"`
	Password       string  `json:"password" validate:"required,min=6" example:"SecurePass123"`
	Role           string  `json:"role" validate:"required,oneof=admin manager sales viewer" example:"sales"`
	TelegramUserID *int64  `json:"telegram_user_id" validate:"omitempty" example:"987654321"`
}

// UserResponse represents the response payload for a user
type UserResponse struct {
	ID             uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Nama           string     `json:"nama" example:"Jane Doe"`
	Email          string     `json:"email" example:"jane.doe@company.com"`
	Role           string     `json:"role" example:"sales"`
	TelegramUserID *int64     `json:"telegram_user_id,omitempty" example:"987654321"`
	Aktif          bool       `json:"aktif" example:"true"`
	CreatedAt      time.Time  `json:"created_at" example:"2026-04-24T10:00:00Z"`
}

// UserListResponse represents the response payload for a list of users
type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

// UserFilter represents the filter parameters for listing users
type UserFilter struct {
	Role   string `query:"role" validate:"omitempty,oneof=admin manager sales viewer"`
	Search string `query:"search" validate:"omitempty"`
	Page   int    `query:"page" validate:"omitempty,min=1" example:"1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100" example:"10"`
}

// UserIDResponse represents the response payload containing only a user ID
type UserIDResponse struct {
	ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Nama           *string `json:"nama" validate:"omitempty,min=2,max=100"`
	Email          *string `json:"email" validate:"omitempty,email"`
	Role           *string `json:"role" validate:"omitempty,oneof=admin manager sales viewer"`
	TelegramUserID *int64  `json:"telegram_user_id" validate:"omitempty"`
}

// DeactivateUserResponse represents the response payload for deactivating a user
type DeactivateUserResponse struct {
	Message string    `json:"message" example:"User deactivated successfully"`
	ID      uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
}