package handler

import (
	"insightflow/be-penjualan/internal/handler/auth"
)

// NewAuthHandler creates a new auth handler instance
func NewAuthHandler() *auth.Handler {
	return auth.New()
}