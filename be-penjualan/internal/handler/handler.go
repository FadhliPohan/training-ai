package handler

import (
	authHandler     "insightflow/be-penjualan/internal/handler/auth"
	customerHandler "insightflow/be-penjualan/internal/handler/customer"
	produkHandler   "insightflow/be-penjualan/internal/handler/produk"
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