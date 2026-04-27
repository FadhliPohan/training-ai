package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/handler"
)

type reportTestEnvelope struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func TestReportEndpoint_ValidRequest(t *testing.T) {
	app := fiber.New()
	reportHandler := handler.NewReportHandler()
	app.Get("/reports", reportHandler.Get)

	req := httptest.NewRequest("GET", "/reports?type=daily-sales&from=2026-04-01&to=2026-04-07", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body reportTestEnvelope
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.True(t, body.Success)
	assert.Equal(t, "daily-sales", body.Data["type"])
	assert.NotEmpty(t, body.Data["summary"])
	assert.NotEmpty(t, body.Data["metrics"])
}

func TestReportEndpoint_InvalidType(t *testing.T) {
	app := fiber.New()
	reportHandler := handler.NewReportHandler()
	app.Get("/reports", reportHandler.Get)

	req := httptest.NewRequest("GET", "/reports?type=unknown", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestReportEndpoint_InvalidMode(t *testing.T) {
	app := fiber.New()
	reportHandler := handler.NewReportHandler()
	app.Get("/reports", reportHandler.Get)

	req := httptest.NewRequest("GET", "/reports?type=daily-sales&mode=foo", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestReportEndpoint_N8NFallbackStillReturns200(t *testing.T) {
	orig := config.App
	defer func() { config.App = orig }()

	// Force n8n to fail quickly so handler should return fallback response.
	config.App.N8NBaseURL = "http://127.0.0.1:1"
	config.App.N8NDashboardWebhookPath = "/webhook/dashboard-ai"
	config.App.N8NWebhookSecret = "test-secret"

	app := fiber.New()
	reportHandler := handler.NewReportHandler()
	app.Get("/reports", reportHandler.Get)

	req := httptest.NewRequest("GET", "/reports?type=daily-sales&from=2026-04-01&to=2026-04-07", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body reportTestEnvelope
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, "fallback", body.Data["ai_source"])
	assert.NotEmpty(t, body.Data["fallback_reason"])
}

func TestReportEndpoint_RawModeBypassN8N(t *testing.T) {
	app := fiber.New()
	reportHandler := handler.NewReportHandler()
	app.Get("/reports", reportHandler.Get)

	req := httptest.NewRequest("GET", "/reports?type=daily-sales&from=2026-04-01&to=2026-04-07&mode=raw", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body reportTestEnvelope
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, "raw", body.Data["ai_source"])
}
