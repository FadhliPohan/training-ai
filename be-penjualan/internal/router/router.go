package router

import (
	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/middleware"
)

// Setup registers all application routes onto the Fiber app instance.
// Route grouping:
//   - Public: /api/v1/auth/*, /api/v1/chat/stream, /api/v1/produk (GET public catalogue)
//   - Auth required: all other routes
//   - Role-guarded: admin-only, manager-only endpoints
func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	// ---- Health Check (public) ----
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "insightflow-api"})
	})

	// ---- Auth (public) ----
	auth := api.Group("/auth")
	{
		// handlers will be registered here once implemented:
		// auth.Post("/login", authHandler.Login)
		// auth.Post("/logout", authHandler.Logout)
		// auth.Post("/register", authHandler.Register)  // customer self-registration
		// auth.Get("/me", middleware.AuthRequired, authHandler.Me)
		_ = auth // placeholder to avoid unused variable error
	}

	// ---- Public Catalogue (no auth needed) ----
	publicProduk := api.Group("/produk")
	{
		// publicProduk.Get("/", produkHandler.ListAktif)
		// publicProduk.Get("/:id", produkHandler.GetByID)
		_ = publicProduk
	}

	// ---- AI Chat Stream (public) ----
	// api.Get("/chat/stream", chatHandler.Stream)

	// ---- Protected routes ----
	protected := api.Group("", middleware.AuthRequired, middleware.ViewerReadOnly)

	// Produk management (admin only for write)
	produk := protected.Group("/produk")
	{
		// produk.Post("/", middleware.RoleGuard("admin"), produkHandler.Create)
		// produk.Put("/:id", middleware.RoleGuard("admin"), produkHandler.Update)
		// produk.Patch("/:id", middleware.RoleGuard("admin"), produkHandler.Deactivate)
		_ = produk
	}

	// Customer management
	customer := protected.Group("/customer")
	{
		// customer.Get("/", customerHandler.List)
		// customer.Get("/:id", customerHandler.GetByID)
		// customer.Post("/", middleware.RoleGuard("admin","sales"), customerHandler.Create)
		// customer.Put("/:id", middleware.RoleGuard("admin","sales"), customerHandler.Update)
		_ = customer
	}

	// User management (admin only)
	users := protected.Group("/users", middleware.RoleGuard("admin"))
	{
		// users.Get("/", userHandler.List)
		// users.Get("/:id", userHandler.GetByID)
		// users.Post("/", userHandler.Create)
		// users.Put("/:id", userHandler.Update)
		// users.Patch("/:id", userHandler.Deactivate)
		_ = users
	}

	// Orders
	orders := protected.Group("/orders")
	{
		// orders.Get("/", orderHandler.List)
		// orders.Get("/:id", orderHandler.GetByID)
		// orders.Post("/", middleware.RoleGuard("admin","sales"), orderHandler.Create)
		// orders.Post("/:id/confirm", middleware.RoleGuard("admin","sales"), orderHandler.Confirm)
		// orders.Post("/:id/cancel", middleware.RoleGuard("admin","sales"), orderHandler.Cancel)
		_ = orders
	}

	// Payments
	payments := protected.Group("/payments")
	{
		// payments.Post("/", middleware.RoleGuard("admin","sales"), paymentHandler.Create)
		// payments.Post("/:id/verify", middleware.RoleGuard("admin"), paymentHandler.Verify)
		_ = payments
	}

	// Shipments
	shipments := protected.Group("/shipments")
	{
		// shipments.Post("/", middleware.RoleGuard("admin","sales"), shipmentHandler.Create)
		// shipments.Put("/:id", middleware.RoleGuard("admin","sales"), shipmentHandler.Update)
		_ = shipments
	}

	// Reports / Dashboard (manager + admin)
	reports := protected.Group("/reports", middleware.RoleGuard("admin","manager","viewer"))
	{
		// reports.Get("/", reportHandler.Get)
		_ = reports
	}

	// Settings (admin only)
	settings := protected.Group("/settings", middleware.RoleGuard("admin"))
	{
		// settings.Get("/telegram", settingHandler.GetTelegram)
		// settings.Put("/telegram", settingHandler.UpdateTelegram)
		_ = settings
	}

	// Internal webhook (called by n8n — protected by webhook secret, not JWT)
	// api.Post("/internal/ai-result", internalHandler.AIResult)
}
