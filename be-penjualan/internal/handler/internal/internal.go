package internal

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/dto"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/response"
	"insightflow/be-penjualan/internal/service"
)

// Handler handles internal endpoints called by n8n workflows.
// All routes in this handler are protected by InternalKeyGuard middleware.
type Handler struct {
	svc         service.InternalService
	reportSvc   service.ReportService
	settingsSvc service.SettingsService
	userRepo    repository.UserRepository
}

// New creates a new internal handler.
func New() *Handler {
	settingsRepo := repository.NewSettingsRepository()
	return &Handler{
		svc:         service.NewInternalService(),
		reportSvc:   service.NewReportService(),
		settingsSvc: service.NewSettingsService(settingsRepo),
		userRepo:    repository.NewUserRepository(),
	}
}

// Summary godoc
//
//	@Summary		Daily sales summary for n8n
//	@Description	Returns flat aggregated sales data for today.
//	@Description	Used by n8n Telegram Daily Summary workflow (scheduled 07:00 WIB).
//	@Description	Auth: X-Internal-Key header required.
//	@Tags			Internal
//	@Produce		json
//	@Param			X-Internal-Key	header		string	true	"Internal API key"
//	@Success		200				{object}	response.Standard
//	@Failure		401				{object}	response.Standard
//	@Failure		500				{object}	response.Standard
//	@Router			/api/internal/reports/summary [get]
func (h *Handler) Summary(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 8*time.Second)
	defer cancel()

	data, err := h.svc.Summary(ctx)
	if err != nil {
		return response.InternalServerError(c)
	}

	return response.OK(c, "Ringkasan laporan harian berhasil diambil", data)
}

// Anomaly godoc
//
//	@Summary		Anomaly detection for n8n alert workflow
//	@Description	Compares today's metrics against a 7-day baseline.
//	@Description	Returns has_anomaly=true and list of anomalies if variance > threshold.
//	@Description	Used by n8n every 15 minutes to decide whether to send Telegram alert.
//	@Description	Auth: X-Internal-Key header required.
//	@Tags			Internal
//	@Produce		json
//	@Param			X-Internal-Key	header		string	true	"Internal API key"
//	@Param			threshold		query		number	false	"Variance threshold % (default: 10)"
//	@Success		200				{object}	response.Standard
//	@Failure		401				{object}	response.Standard
//	@Failure		500				{object}	response.Standard
//	@Router			/api/internal/reports/anomaly [get]
func (h *Handler) Anomaly(c *fiber.Ctx) error {
	thresholdStr := c.Query("threshold", "10")
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil || threshold <= 0 {
		threshold = 10.0
	}

	ctx, cancel := context.WithTimeout(c.Context(), 8*time.Second)
	defer cancel()

	data, err := h.svc.Anomaly(ctx, threshold)
	if err != nil {
		return response.InternalServerError(c)
	}

	return response.OK(c, "Pengecekan anomali selesai", data)
}

// UserByTelegram godoc
//
//	@Summary		Resolve user by Telegram user ID
//	@Description	Returns user profile (id, nama, role) for a given telegram_user_id.
//	@Description	Used by n8n Telegram Q&A workflow to verify sender and determine data scope.
//	@Description	Auth: X-Internal-Key header required.
//	@Tags			Internal
//	@Produce		json
//	@Param			X-Internal-Key		header		string	true	"Internal API key"
//	@Param			telegram_user_id	query		integer	true	"Telegram user ID (int64)"
//	@Success		200					{object}	response.Standard
//	@Failure		400					{object}	response.Standard
//	@Failure		401					{object}	response.Standard
//	@Failure		404					{object}	response.Standard
//	@Router			/api/internal/users/by-telegram [get]
func (h *Handler) UserByTelegram(c *fiber.Ctx) error {
	tidStr := c.Query("telegram_user_id")
	if tidStr == "" {
		return response.BadRequest(c, "Parameter telegram_user_id wajib diisi.", nil)
	}

	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil || tid <= 0 {
		return response.BadRequest(c, "telegram_user_id harus berupa angka positif.", nil)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	user, err := h.userRepo.FindByTelegramUserID(ctx, tid)
	if err != nil {
		return response.NotFound(c, "Pengguna dengan Telegram ID tersebut tidak ditemukan atau tidak aktif.")
	}

	return response.OK(c, "Data pengguna ditemukan", fiber.Map{
		"user_id": user.ID.String(),
		"nama":    user.Nama,
		"role":    user.Role,
	})
}

// Reports godoc
//
//	@Summary		Dynamic report data for n8n Telegram Q&A
//	@Description	Returns raw report data for a given type — used by n8n Telegram Q&A workflow.
//	@Description	n8n parses user intent, picks the type, calls this endpoint, then passes data to LLM.
//	@Description	Always returns mode=raw (no AI enrichment) — n8n handles the LLM call itself.
//	@Description	Role scoping: if sales_id is provided, only that sales' data is returned.
//	@Description	Auth: X-Internal-Key header required.
//	@Tags			Internal
//	@Produce		json
//	@Param			X-Internal-Key	header		string	true	"Internal API key"
//	@Param			type			query		string	true	"Report type: daily-sales|monthly-sales|top-products|sales-by-person|order-funnel|category-breakdown|low-stock|revenue-trend"
//	@Param			from			query		string	false	"From date (YYYY-MM-DD)"
//	@Param			to				query		string	false	"To date (YYYY-MM-DD)"
//	@Param			sales_id		query		string	false	"Sales UUID — if set, scopes data to that sales only"
//	@Success		200				{object}	response.Standard
//	@Failure		400				{object}	response.Standard
//	@Failure		401				{object}	response.Standard
//	@Failure		500				{object}	response.Standard
//	@Router			/api/internal/reports [get]
func (h *Handler) Reports(c *fiber.Ctx) error {
	reportType := strings.TrimSpace(c.Query("type"))
	if reportType == "" {
		return response.BadRequest(c, "Parameter type wajib diisi. Pilihan: daily-sales, top-products, low-stock, order-funnel, category-breakdown, sales-by-person, monthly-sales, revenue-trend.", nil)
	}

	req := dto.ReportRequest{
		Type:    reportType,
		From:    strings.TrimSpace(c.Query("from")),
		To:      strings.TrimSpace(c.Query("to")),
		SalesID: strings.TrimSpace(c.Query("sales_id")),
		Mode:    "raw", // always raw — n8n handles LLM call itself
	}

	ctx, cancel := context.WithTimeout(c.Context(), 8*time.Second)
	defer cancel()

	result, err := h.reportSvc.Generate(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "jenis report tidak valid") {
			return response.BadRequest(c, err.Error(), nil)
		}
		if strings.Contains(err.Error(), "format tanggal") || strings.Contains(err.Error(), "rentang tanggal") {
			return response.BadRequest(c, err.Error(), nil)
		}
		return response.InternalServerError(c)
	}

	return response.OK(c, "Data laporan berhasil diambil", result)
}

// Settings godoc
//
//	@Summary		Get Telegram settings for n8n
//	@Description	Returns active Telegram config (chat_id, jam_summary) and anomaly threshold.
//	@Description	n8n uses this to dynamically get chat_id and threshold — no hardcoding needed.
//	@Description	Auth: X-Internal-Key header required.
//	@Tags			Internal
//	@Produce		json
//	@Param			X-Internal-Key	header		string	true	"Internal API key"
//	@Success		200				{object}	response.Standard
//	@Failure		401				{object}	response.Standard
//	@Failure		500				{object}	response.Standard
//	@Router			/api/internal/settings/telegram [get]
func (h *Handler) Settings(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	result, err := h.settingsSvc.GetTelegram(ctx)
	if err != nil {
		return response.InternalServerError(c)
	}

	return response.OK(c, "Konfigurasi Telegram berhasil diambil", result)
}
