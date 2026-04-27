package settings

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/dto"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/response"
	"insightflow/be-penjualan/internal/service"
)

// Handler handles settings endpoints.
type Handler struct {
	svc service.SettingsService
}

// New creates settings handler instance.
func New() *Handler {
	repo := repository.NewSettingsRepository()
	svc := service.NewSettingsService(repo)
	return &Handler{svc: svc}
}

// GetTelegram handles GET /api/v1/settings/telegram.
//
//	@Summary		Get telegram settings
//	@Description	Get current telegram config and anomaly threshold
//	@Tags			Settings
//	@Security		JWT
//	@Produce		json
//	@Success		200	{object}	response.Standard
//	@Failure		500	{object}	response.Standard
//	@Router			/settings/telegram [get]
func (h *Handler) GetTelegram(c *fiber.Ctx) error {
	result, err := h.svc.GetTelegram(c.Context())
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.OK(c, "Berhasil mendapatkan konfigurasi Telegram", result)
}

// UpdateTelegram handles PUT /api/v1/settings/telegram.
//
//	@Summary		Update telegram settings
//	@Description	Update telegram config and anomaly threshold
//	@Tags			Settings
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UpdateTelegramConfigRequest	true	"Telegram settings payload"
//	@Success		200		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		500		{object}	response.Standard
//	@Router			/settings/telegram [put]
func (h *Handler) UpdateTelegram(c *fiber.Ctx) error {
	var req dto.UpdateTelegramConfigRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}
	if req.NamaGrup != nil && strings.TrimSpace(*req.NamaGrup) == "" {
		return response.BadRequest(c, "nama_grup tidak boleh kosong", nil)
	}
	if req.JamSummary != nil {
		if _, err := time.Parse("15:04", strings.TrimSpace(*req.JamSummary)); err != nil {
			return response.BadRequest(c, service.ErrInvalidJamSummary.Error(), nil)
		}
	}
	if req.ThresholdPct != nil && (*req.ThresholdPct < 0 || *req.ThresholdPct > 100) {
		return response.BadRequest(c, service.ErrInvalidThreshold.Error(), nil)
	}

	result, err := h.svc.UpdateTelegram(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidJamSummary) || errors.Is(err, service.ErrInvalidThreshold) {
			return response.BadRequest(c, err.Error(), nil)
		}
		if err.Error() == "nama_grup tidak boleh kosong" {
			return response.BadRequest(c, err.Error(), nil)
		}
		return response.InternalServerError(c)
	}
	return response.OK(c, "Konfigurasi Telegram berhasil diupdate", result)
}
