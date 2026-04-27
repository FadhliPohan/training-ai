package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"insightflow/be-penjualan/internal/handler"
	"insightflow/be-penjualan/internal/middleware"
)

// Setup registers all application routes onto the Fiber app instance.
// Route grouping:
//   - Public: /api/v1/auth/*, /api/v1/produk (GET public catalogue)
//   - Auth required: all other routes
//   - Role-guarded: admin-only, manager-only endpoints
func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	// ---- Health Check (public) ----
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "insightflow-api"})
	})

	// ---- Swagger UI ----
	app.Get("/swagger/*", swagger.HandlerDefault)

	// ---- Auth (public + protected /me) ----
	auth := api.Group("/auth")
	{
		authHandler := handler.NewAuthHandler()
		authHandler.RegisterRoutes(auth)
	}

	// ---- Handlers ----
	produkHandler := handler.NewProdukHandler()
	customerHandler := handler.NewCustomerHandler()
	usersHandler := handler.NewUsersHandler()

	// ---- Public Catalogue (no auth needed) ----
	// GET /api/v1/produk and /api/v1/produk/:id accessible without token
	api.Get("/produk", produkHandler.ListPublic)
	api.Get("/produk/:id", produkHandler.GetByID)

	// ---- Protected routes ----
	protected := api.Group("", middleware.AuthRequired, middleware.ViewerReadOnly)

	// Produk management — write operations (admin only)
	produk := protected.Group("/produk")
	{
		produk.Post("/", middleware.RoleGuard("admin"), produkHandler.Create)
		produk.Put("/:id", middleware.RoleGuard("admin"), produkHandler.Update)
		produk.Patch("/:id", middleware.RoleGuard("admin"), produkHandler.Deactivate)
	}

	// Customer management (admin + sales for write, all auth for read)
	customer := protected.Group("/customer")
	{
		customer.Get("/", customerHandler.List)
		customer.Get("/:id", customerHandler.GetByID)
		customer.Post("/", middleware.RoleGuard("admin", "sales"), customerHandler.Create)
		customer.Put("/:id", middleware.RoleGuard("admin", "sales"), customerHandler.Update)
	}

	// User management (admin only)
	users := protected.Group("/users", middleware.RoleGuard("admin"))
	{
		users.Get("/", usersHandler.List)
		users.Get("/:id", usersHandler.GetByID)
		users.Post("/", usersHandler.Create)
		users.Put("/:id", usersHandler.Update)
		users.Patch("/:id", usersHandler.Deactivate)
	}

	// Orders — Sprint 2
	orders := protected.Group("/orders")
	{
		_ = orders
	}

	// Payments — Sprint 2
	payments := protected.Group("/payments")
	{
		_ = payments
	}

	// Shipments — Sprint 2
	shipments := protected.Group("/shipments")
	{
		_ = shipments
	}

	// Reports / Dashboard — Sprint 3
	reports := protected.Group("/reports", middleware.RoleGuard("admin", "manager", "viewer"))
	{
		_ = reports
	}

	// Settings — Sprint 3
	settings := protected.Group("/settings", middleware.RoleGuard("admin"))
	{
		_ = settings
	}

	// Internal webhook (called by n8n — Sprint 3)
	// api.Post("/internal/ai-result", internalHandler.AIResult)
}
