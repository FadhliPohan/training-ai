//	@title			Auth API
//	@version		1.0
//	@description	Authentication endpoints

package auth

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/database"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/response"
)

// LoginRequest represents the expected payload for the login endpoint.
//
//	@Description	User login request
type LoginRequest struct {
	// User email address
	Email string `json:"email" validate:"required,email" example:"admin@insightflow.id"`
	// User password (min 6 characters)
	Password string `json:"password" validate:"required,min=6" example:"Admin@12345"`
}

// LoginResponse represents the successful login response with JWT token.
//
//	@Description	User login response
type LoginResponse struct {
	// JWT token for authentication
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// Token expiration time
	ExpiresAt time.Time `json:"expires_at" example:"2026-04-25T01:18:10Z"`
	// User information
	User UserInfo `json:"user"`
}

// UserInfo contains non-sensitive user information returned after login.
//
//	@Description	User information
type UserInfo struct {
	// User unique identifier
	ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	// User full name
	Nama string `json:"nama" example:"Administrator"`
	// User email address
	Email string `json:"email" example:"admin@insightflow.id"`
	// User role (admin, manager, sales, viewer)
	Role string `json:"role" example:"admin"`
}

// Login handles POST /api/v1/auth/login
// Validates user credentials and returns a JWT token on success.
//
//	@Summary		User login
//	@Description	Authenticate user with email and password to get JWT token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest	true	"Login credentials"
//	@Success		200		{object}	response.Standard{data=LoginResponse}
//	@Failure		400		{object}	response.Standard
//	@Failure		401		{object}	response.Standard
//	@Failure		500		{object}	response.Standard
//	@Router			/auth/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	// Parse request body
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		return response.BadRequest(c, "Email dan password harus diisi", nil)
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return response.BadRequest(c, "Format email tidak valid", nil)
	}

	// Find user by email
	user, err := findUserByEmail(req.Email)
	if err != nil {
		// For security, we don't reveal if user exists or not
		return response.Unauthorized(c, "Email atau password salah")
	}

	// Check if user is active
	if !user.Aktif {
		return response.Unauthorized(c, "Akun tidak aktif. Silakan hubungi administrator.")
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return response.Unauthorized(c, "Email atau password salah")
	}

	// Generate JWT token
	token, expiresAt, err := generateJWTToken(user)
	if err != nil {
		return response.InternalServerError(c)
	}

	// Prepare response
	resp := LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User: UserInfo{
			ID:    user.ID,
			Nama:  user.Nama,
			Email: user.Email,
			Role:  user.Role,
		},
	}

	return response.OK(c, "Login berhasil", resp)
}

// findUserByEmail queries the database for a user with the given email.
func findUserByEmail(email string) (*domain.User, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool is not initialized")
	}

	query := `
		SELECT id, nama, email, password, role, telegram_user_id, aktif, created_at
		FROM app.users
		WHERE email = $1
	`

	var user domain.User
	err := database.Pool.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Nama,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.TelegramUserID,
		&user.Aktif,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// generateJWTToken creates a signed JWT token for the given user.
func generateJWTToken(user *domain.User) (string, time.Time, error) {
	// Token expires in 8 hours
	expiresAt := time.Now().Add(8 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
		"exp":     expiresAt.Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.App.JWTSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expiresAt, nil
}
