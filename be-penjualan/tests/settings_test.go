package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"insightflow/be-penjualan/internal/handler"
)

func TestUpdateTelegramSettings_InvalidJamSummary(t *testing.T) {
	app := fiber.New()
	settingsHandler := handler.NewSettingsHandler()
	app.Put("/settings/telegram", settingsHandler.UpdateTelegram)

	payload := map[string]interface{}{
		"jam_summary": "25:99",
	}
	raw, _ := json.Marshal(payload)

	req := httptest.NewRequest("PUT", "/settings/telegram", bytes.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestUpdateTelegramSettings_InvalidThreshold(t *testing.T) {
	app := fiber.New()
	settingsHandler := handler.NewSettingsHandler()
	app.Put("/settings/telegram", settingsHandler.UpdateTelegram)

	payload := map[string]interface{}{
		"threshold_pct": 120,
	}
	raw, _ := json.Marshal(payload)

	req := httptest.NewRequest("PUT", "/settings/telegram", bytes.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
