package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"insightflow/be-penjualan/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Initialize database connection for tests
	// In a real test environment, you would use a test database
	// For now, we'll skip database initialization in tests
	// database.Connect()
}

// TestLoginEndpoint tests the POST /api/v1/auth/login endpoint
func TestLoginEndpoint(t *testing.T) {
	// Setup
	app := fiber.New()
	router.Setup(app)

	// Test cases
	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
	}{
		{
			name: "Invalid credentials",
			payload: map[string]interface{}{
				"email":    "admin@insightflow.id",
				"password": "wrongpassword",
			},
			expectedStatus: 401,
		},
		{
			name: "Missing email",
			payload: map[string]interface{}{
				"password": "Admin@12345",
			},
			expectedStatus: 400,
		},
		{
			name: "Missing password",
			payload: map[string]interface{}{
				"email": "admin@insightflow.id",
			},
			expectedStatus: 400,
		},
		{
			name: "Invalid email format",
			payload: map[string]interface{}{
				"email":    "invalid-email",
				"password": "Admin@12345",
			},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert payload to JSON
			payloadBytes, _ := json.Marshal(tt.payload)

			// Create request
			req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewReader(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			resp, err := app.Test(req, -1) // -1 means no timeout
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
