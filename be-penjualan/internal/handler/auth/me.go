package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"insightflow/be-penjualan/internal/middleware"
	"insightflow/be-penjualan/internal/response"
)

// Me handles GET /api/v1/auth/me — returns the current authenticated user profile.
//
//	@Summary		Get current user
//	@Description	Returns profile of the currently authenticated user
//	@Tags			Auth
//	@Security		JWT
//	@Produce		json
//	@Success		200	{object}	response.Standard
//	@Failure		401	{object}	response.Standard
//	@Router			/auth/me [get]
func (h *Handler) Me(c *fiber.Ctx) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return response.Unauthorized(c, "")
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return response.Unauthorized(c, "Token tidak valid")
	}

	user, err := h.svc.GetUserByID(c.Context(), userID)
	if err != nil {
		return response.NotFound(c, "User tidak ditemukan")
	}

	return response.OK(c, "Berhasil mendapatkan data user", fiber.Map{
		"id":               user.ID,
		"nama":             user.Nama,
		"email":            user.Email,
		"role":             user.Role,
		"telegram_user_id": user.TelegramUserID,
		"aktif":            user.Aktif,
		"created_at":       user.CreatedAt,
	})
}
