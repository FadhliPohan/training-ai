package auth

import (
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/response"
)

// RegisterRequest represents the payload for customer self-registration.
//
//	@Description	Customer self-registration request
type RegisterRequest struct {
	Nama     string `json:"nama" validate:"required" example:"Budi Santoso"`
	Email    string `json:"email" validate:"required,email" example:"budi@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"Budi@12345"`
}

// Register handles POST /api/v1/auth/register — customer self-registration.
//
//	@Summary		Customer self-registration
//	@Description	Register a new customer account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RegisterRequest	true	"Registration data"
//	@Success		201		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		409		{object}	response.Standard
//	@Router			/auth/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	// Validasi input
	if req.Nama == "" || req.Email == "" || req.Password == "" {
		return response.BadRequest(c, "Nama, email, dan password harus diisi", nil)
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return response.BadRequest(c, "Format email tidak valid", nil)
	}
	if len(req.Password) < 8 {
		return response.BadRequest(c, "Password minimal 8 karakter", nil)
	}

	// Register via service (role default: "viewer" for self-registration)
	user, err := h.svc.Register(c.Context(), req.Nama, req.Email, req.Password, "viewer", nil)
	if err != nil {
		if err.Error() == "email sudah terdaftar" {
			return response.Conflict(c, err.Error())
		}
		return response.InternalServerError(c)
	}

	return response.Created(c, "Registrasi berhasil", fiber.Map{
		"id":    user.ID,
		"nama":  user.Nama,
		"email": user.Email,
		"role":  user.Role,
	})
}
