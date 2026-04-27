package auth

import (
	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/middleware"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/service"
)

// Handler holds auth service dependency.
type Handler struct {
	svc service.AuthService
}

// New creates a new auth Handler wired up with real repository/service.
func New() *Handler {
	repo := repository.NewUserRepository()
	svc := service.NewAuthService(repo)
	return &Handler{svc: svc}
}

// RegisterRoutes registers all auth-related routes.
func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/login", h.Login)
	router.Post("/logout", h.Logout)
	router.Post("/register", h.Register)
	router.Get("/me", middleware.AuthRequired, h.Me)
}