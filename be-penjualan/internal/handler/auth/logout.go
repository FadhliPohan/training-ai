package auth

import (
	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/response"
)

// Logout handles POST /api/v1/auth/logout
// Clears the access_token cookie. Frontend is responsible for removing Bearer token from memory.
//
//	@Summary		Logout
//	@Description	Logout the current user (clears httpOnly cookie)
//	@Tags			Auth
//	@Security		JWT
//	@Produce		json
//	@Success		200	{object}	response.Standard
//	@Router			/auth/logout [post]
func (h *Handler) Logout(c *fiber.Ctx) error {
	// Clear the httpOnly cookie if it was used
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		SameSite: "Lax",
	})
	return response.OK(c, "Berhasil logout", nil)
}
