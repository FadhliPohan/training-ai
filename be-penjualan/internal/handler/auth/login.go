package auth

import (
	"net/mail"

	"github.com/gofiber/fiber/v2"
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

// Login handles POST /api/v1/auth/login
//
//	@Summary		User login
//	@Description	Authenticate user with email and password to get JWT token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest	true	"Login credentials"
//	@Success		200		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		401		{object}	response.Standard
//	@Router			/auth/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}
	if req.Email == "" || req.Password == "" {
		return response.BadRequest(c, "Email dan password harus diisi", nil)
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return response.BadRequest(c, "Format email tidak valid", nil)
	}

	user, token, expiresAt, err := h.svc.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	return response.OK(c, "Login berhasil", fiber.Map{
		"token":      token,
		"expires_at": expiresAt,
		"user": fiber.Map{
			"id":    user.ID,
			"nama":  user.Nama,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
