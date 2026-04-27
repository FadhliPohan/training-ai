package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== ORDER DTOs ====================

// OrderDetailRequest represents the request payload for an order detail item
type OrderDetailRequest struct {
	ProdukID int `json:"produk_id" validate:"required" example:"1"`
	Qty      int `json:"qty" validate:"required,min=1" example:"2"`
}

// CreateOrderRequest represents the request payload for creating an order
type CreateOrderRequest struct {
	CustomerID int                  `json:"customer_id" validate:"required" example:"1"`
	Details    []OrderDetailRequest `json:"details" validate:"required,min=1"`
}

// OrderDetailResponse represents the response payload for an order detail
type OrderDetailResponse struct {
	ID        int     `json:"id" example:"1"`
	ProdukID  int     `json:"produk_id" example:"1"`
	Qty       int     `json:"qty" example:"2"`
	HargaSaat float64 `json:"harga_saat" example:"89000"`
	Subtotal  float64 `json:"subtotal" example:"178000"`
	Produk    ProdukResponse `json:"produk,omitempty"`
}

// OrderResponse represents the response payload for an order
type OrderResponse struct {
	ID         int                   `json:"id" example:"1"`
	NoOrder    string                `json:"no_order" example:"ORD-20260424-001"`
	CustomerID int                   `json:"customer_id" example:"1"`
	SalesID    uuid.UUID             `json:"sales_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Tanggal    time.Time             `json:"tanggal" example:"2026-04-24T10:00:00Z"`
	Status     string                `json:"status" example:"pending"`
	Total      float64               `json:"total" example:"178000"`
	CreatedAt  time.Time             `json:"created_at" example:"2026-04-24T10:00:00Z"`
	Customer   *CustomerResponse     `json:"customer,omitempty"`
	Sales      *UserResponse         `json:"sales,omitempty"`
	Details    []OrderDetailResponse `json:"details,omitempty"`
}

// OrderListResponse represents the response payload for a list of orders
type OrderListResponse struct {
	Orders []OrderResponse `json:"orders"`
}

// OrderFilter represents the filter parameters for listing orders
type OrderFilter struct {
	Status     string    `query:"status" validate:"omitempty,oneof=pending confirmed paid shipped closed cancelled"`
	CustomerID *int      `query:"customer_id" validate:"omitempty"`
	SalesID    *string   `query:"sales_id" validate:"omitempty,uuid"`
	FromDate   time.Time `query:"from_date" validate:"omitempty"`
	ToDate     time.Time `query:"to_date" validate:"omitempty"`
	Page       int       `query:"page" validate:"omitempty,min=1" example:"1"`
	Limit      int       `query:"limit" validate:"omitempty,min=1,max=100" example:"10"`
}

// OrderIDResponse represents the response payload containing only an order ID
type OrderIDResponse struct {
	ID int `json:"id" example:"1"`
}

// ConfirmOrderRequest represents the request payload for confirming an order
type ConfirmOrderRequest struct {
	// No additional fields needed for confirmation
}

// CancelOrderRequest represents the request payload for cancelling an order
type CancelOrderRequest struct {
	Alasan string `json:"alasan" validate:"required,min=1,max=500" example:"Customer requested cancellation"`
}

// OrderStatusResponse represents the response payload for order status changes
type OrderStatusResponse struct {
	ID     int    `json:"id" example:"1"`
	Status string `json:"status" example:"confirmed"`
	Message string `json:"message" example:"Order confirmed successfully"`
}