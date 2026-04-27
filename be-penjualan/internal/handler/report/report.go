package report

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/dto"
	"insightflow/be-penjualan/internal/response"
	"insightflow/be-penjualan/internal/service"
)

// Handler handles dashboard report endpoints.
type Handler struct {
	svc service.ReportService
}

// New creates a report handler instance.
func New() *Handler {
	return &Handler{
		svc: service.NewReportService(),
	}
}

// Get handles GET /api/v1/reports.
//
//	@Summary		Get dashboard report summary (dummy data + AI enrichment)
//	@Description	Return aggregated report for dashboard and enrich with n8n/AI when available
//	@Tags			Reports
//	@Security		JWT
//	@Produce		json
//	@Param			type		query		string	false	"Report type"
//	@Param			from		query		string	false	"From date (YYYY-MM-DD or RFC3339)"
//	@Param			to			query		string	false	"To date (YYYY-MM-DD or RFC3339)"
//	@Param			sales_id	query		string	false	"Sales UUID"
//	@Param			mode		query		string	false	"Mode: ai (default) | raw"
//	@Success		200			{object}	response.Standard
//	@Failure		400			{object}	response.Standard
//	@Failure		500			{object}	response.Standard
//	@Router			/reports [get]
func (h *Handler) Get(c *fiber.Ctx) error {
	req := dto.ReportRequest{
		Type:    strings.TrimSpace(c.Query("type")),
		From:    strings.TrimSpace(c.Query("from")),
		To:      strings.TrimSpace(c.Query("to")),
		SalesID: strings.TrimSpace(c.Query("sales_id")),
		Mode:    strings.TrimSpace(c.Query("mode")),
	}

	result, err := h.svc.Generate(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidReportType) ||
			errors.Is(err, service.ErrInvalidDateFormat) ||
			errors.Is(err, service.ErrInvalidDateRange) ||
			errors.Is(err, service.ErrInvalidReportMode) {
			return response.BadRequest(c, err.Error(), nil)
		}
		return response.InternalServerError(c)
	}

	return response.OK(c, "Berhasil mendapatkan ringkasan laporan", result)
}
