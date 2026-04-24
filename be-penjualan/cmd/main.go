package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/database"
	"insightflow/be-penjualan/internal/router"
)

func main() {
	// ---- Logging ----
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("APP_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}

	// ---- Config ----
	config.Load()
	log.Info().Str("env", config.App.AppEnv).Msg("configuration loaded")

	// ---- Database ----
	database.Connect()
	defer database.Close()

	// ---- Fiber App ----
	app := fiber.New(fiber.Config{
		AppName:      "InsightFlow API v1",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		// Do not expose internal error details to clients
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Error().Err(err).Str("path", c.Path()).Int("status", code).Msg("request error")
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": "Terjadi kesalahan pada server. Silakan coba beberapa saat lagi.",
			})
		},
	})

	// ---- Global Middleware ----
	app.Use(recover.New()) // catch panics

	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} ${method} ${path} ${latency}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.App.FrontendURL,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true, // required for httpOnly cookie auth
	}))

	// ---- Routes ----
	router.Setup(app)

	// ---- Start Server (with graceful shutdown) ----
	addr := fmt.Sprintf(":%s", config.App.AppPort)
	log.Info().Str("address", addr).Msg("starting InsightFlow API server")

	// Listen in a goroutine so we can handle shutdown signals
	go func() {
		if err := app.Listen(addr); err != nil {
			log.Fatal().Err(err).Msg("server listen failed")
		}
	}()

	// Wait for interrupt signal (Ctrl+C / SIGTERM from Docker/systemd)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutdown signal received, gracefully stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Error().Err(err).Msg("error during graceful shutdown")
	}

	log.Info().Msg("server stopped")
}
