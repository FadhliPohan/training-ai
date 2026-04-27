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
	reportHandler := handler.NewReportHandler()
	settingsHandler := handler.NewSettingsHandler()
	chatHandler := handler.NewChatHandler()

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
		reports.Get("/", reportHandler.Get)
	}

	// AI Chat — OpenAI GPT-4o (all authenticated roles)
	protected.Post("/chat", chatHandler.Chat)

	// Settings — Sprint 3
	settings := protected.Group("/settings", middleware.RoleGuard("admin"))
	{
		settings.Get("/telegram", settingsHandler.GetTelegram)
		settings.Put("/telegram", settingsHandler.UpdateTelegram)
	}

	// ---- Telegram inbound webhook (public — called by Telegram servers) ----
	// Telegram sends POST requests to this URL whenever a user messages the bot.
	// No JWT required — Telegram does not support custom auth headers on its side.
	// Security: validate X-Telegram-Bot-Api-Secret-Token inside handler (optional).
	telegramHandler := handler.NewTelegramHandler()
	api.Post("/telegram/webhook", telegramHandler.Webhook)

	// ---- Internal routes (called by n8n — protected by X-Internal-Key) ----
	// These endpoints bypass JWT and are meant only for internal service-to-service calls.
	internalHandler := handler.NewInternalHandler()
	internal := app.Group("/api/internal", middleware.InternalKeyGuard)
	{
		// GET /api/internal/reports/summary  → n8n Daily Summary workflow
		// GET /api/internal/reports/anomaly  → n8n Anomaly Alert workflow (every 15 min)
		// GET /api/internal/reports          → n8n Telegram Q&A (dynamic intent routing)
		internalReports := internal.Group("/reports")
		internalReports.Get("/summary", internalHandler.Summary)
		internalReports.Get("/anomaly", internalHandler.Anomaly)
		internalReports.Get("/", internalHandler.Reports)

		// GET /api/internal/users/by-telegram → n8n Telegram Q&A workflow (user lookup)
		internalUsers := internal.Group("/users")
		internalUsers.Get("/by-telegram", internalHandler.UserByTelegram)

		// GET /api/internal/settings/telegram → n8n reads chat_id + threshold dynamically
		// Admin updates these via PUT /api/v1/settings/telegram (JWT-protected)
		// n8n picks them up here without needing JWT
		internalSettings := internal.Group("/settings")
		internalSettings.Get("/telegram", internalHandler.Settings)
	}
}
