package auth

import (
	"github.com/gofiber/fiber/v2"
)

// Handler struct holds dependencies for auth handlers
type Handler struct{}

// New creates a new auth handler
func New() *Handler {
	return &Handler{}
}

// RegisterRoutes registers all auth-related routes
func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/login", h.Login)
	// router.Post("/logout", h.Logout)
	// router.Post("/register", h.Register)
	// router.Get("/me", h.Me)
}