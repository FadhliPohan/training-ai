package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all application configuration values loaded from environment variables.
type Config struct {
	AppEnv                  string
	AppPort                 string
	DatabaseURL             string
	JWTSecret               string
	AutoMigrate             bool
	InternalAPIKey          string
	N8NBaseURL              string
	N8NWebhookSecret        string
	N8NDashboardWebhookPath string
	N8NChatWebhookPath      string
	N8NTelegramWebhookPath  string
	FrontendURL             string
	TelegramBotToken        string
	OpenAIAPIKey            string
	OpenAIURL               string
}

var App Config

// Load reads environment variables and populates the App config.
// It loads from .env file if present (development), or uses OS env (production).
func Load() {
	// Load .env file if it exists (non-fatal if missing in production)
	if err := godotenv.Load(); err != nil {
		log.Println("[config] .env file not found, using OS environment variables")
	}

	App = Config{
		AppEnv:                  getEnv("APP_ENV", "development"),
		AppPort:                 getEnv("APP_PORT", "8080"),
		DatabaseURL:             mustEnv("DATABASE_URL"),
		JWTSecret:               mustEnv("JWT_SECRET"),
		AutoMigrate:             false,
		InternalAPIKey:          mustEnv("INTERNAL_API_KEY"),
		N8NBaseURL:              getEnv("N8N_BASE_URL", "http://localhost:5678"),
		N8NWebhookSecret:        getEnv("N8N_WEBHOOK_SECRET", ""),
		N8NDashboardWebhookPath: getEnv("N8N_DASHBOARD_WEBHOOK_PATH", "/webhook/dashboard-ai"),
		N8NChatWebhookPath:      getEnv("N8N_CHAT_WEBHOOK_PATH", "/webhook/chat-ai"),
		N8NTelegramWebhookPath:  getEnv("N8N_TELEGRAM_WEBHOOK_PATH", "/webhook/telegram-qa"),
		FrontendURL:             getEnv("FRONTEND_URL", "http://localhost:3000"),
		TelegramBotToken:        getEnv("TELEGRAM_BOT_TOKEN", ""),
		OpenAIAPIKey:            getEnv("OPENAI_API_KEY", ""),
		OpenAIURL:               getEnv("OPENAI_URL", "https://api.openai.com/v1/chat/completions"),
	}

	// Default behavior:
	// - production/staging: run auto-migration
	// - development: skip auto-migration
	autoMigrateDefault := App.AppEnv == "production" || App.AppEnv == "staging"
	App.AutoMigrate = getEnvBool("AUTO_MIGRATE", autoMigrateDefault)

	validateJWTSecret(App.JWTSecret)
}

// getEnv returns the value of the environment variable or a fallback.
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// mustEnv panics if the required environment variable is not set.
func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("[config] required environment variable %q is not set", key)
	}
	return val
}

// getEnvInt returns the integer value of the environment variable or a fallback.
func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("[config] invalid integer value for %q, using fallback %d", key, fallback)
		return fallback
	}
	return n
}

// getEnvBool returns boolean env value or fallback.
func getEnvBool(key string, fallback bool) bool {
	val := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	if val == "" {
		return fallback
	}

	switch val {
	case "1", "true", "yes", "y", "on":
		return true
	case "0", "false", "no", "n", "off":
		return false
	default:
		log.Printf("[config] invalid boolean value for %q: %q, using fallback %v", key, val, fallback)
		return fallback
	}
}

// validateJWTSecret enforces a minimum key length of 32 characters.
func validateJWTSecret(secret string) {
	if len(secret) < 32 {
		log.Fatal("[config] JWT_SECRET must be at least 32 characters long")
	}
}

// IsDevelopment returns true when running in development mode.
func IsDevelopment() bool {
	return App.AppEnv == "development"
}
