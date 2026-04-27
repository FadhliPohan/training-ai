package handler

import (
	authHandler     "insightflow/be-penjualan/internal/handler/auth"
	customerHandler "insightflow/be-penjualan/internal/handler/customer"
	internalHandler "insightflow/be-penjualan/internal/handler/internal"
	produkHandler   "insightflow/be-penjualan/internal/handler/produk"
	reportHandler   "insightflow/be-penjualan/internal/handler/report"
	settingsHandler "insightflow/be-penjualan/internal/handler/settings"
	telegramHandler "insightflow/be-penjualan/internal/handler/telegram"
	usersHandler    "insightflow/be-penjualan/internal/handler/users"
)

// NewAuthHandler creates a new auth handler instance.
func NewAuthHandler() *authHandler.Handler {
	return authHandler.New()
}

// NewProdukHandler creates a new produk handler instance.
func NewProdukHandler() *produkHandler.Handler {
	return produkHandler.New()
}

// NewCustomerHandler creates a new customer handler instance.
func NewCustomerHandler() *customerHandler.Handler {
	return customerHandler.New()
}

// NewUsersHandler creates a new users handler instance.
func NewUsersHandler() *usersHandler.Handler {
	return usersHandler.New()
}

// NewReportHandler creates a new report handler instance.
func NewReportHandler() *reportHandler.Handler {
	return reportHandler.New()
}

// NewSettingsHandler creates a new settings handler instance.
func NewSettingsHandler() *settingsHandler.Handler {
	return settingsHandler.New()
}

// NewTelegramHandler creates a new Telegram webhook handler instance.
func NewTelegramHandler() *telegramHandler.Handler {
	return telegramHandler.New()
}

// NewInternalHandler creates a new internal (n8n-facing) handler instance.
func NewInternalHandler() *internalHandler.Handler {
	return internalHandler.New()
}
