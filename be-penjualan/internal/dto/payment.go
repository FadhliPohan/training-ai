package dto

import "time"

// ==================== PAYMENT DTOs ====================

// CreatePaymentRequest represents the request payload for creating a payment
type CreatePaymentRequest struct {
	OrderID int     `json:"order_id" validate:"required" example:"1"`
	Jumlah  float64 `json:"jumlah" validate:"required,min=0" example:"178000"`
	Metode  string  `json:"metode" validate:"required,oneof=transfer tunai kartu" example:"transfer"`
}

// PaymentResponse represents the response payload for a payment
type PaymentResponse struct {
	ID      int       `json:"id" example:"1"`
	OrderID int       `json:"order_id" example:"1"`
	Jumlah  float64   `json:"jumlah" example:"178000"`
	Metode  string    `json:"metode" example:"transfer"`
	Status  string    `json:"status" example:"verified"`
	Tanggal time.Time `json:"tanggal" example:"2026-04-24T10:30:00Z"`
}

// PaymentListResponse represents the response payload for a list of payments
type PaymentListResponse struct {
	Payments []PaymentResponse `json:"payments"`
}

// PaymentFilter represents the filter parameters for listing payments
type PaymentFilter struct {
	Status   string    `query:"status" validate:"omitempty,oneof=pending verified rejected"`
	OrderID  *int      `query:"order_id" validate:"omitempty"`
	FromDate time.Time `query:"from_date" validate:"omitempty"`
	ToDate   time.Time `query:"to_date" validate:"omitempty"`
	Page     int       `query:"page" validate:"omitempty,min=1" example:"1"`
	Limit    int       `query:"limit" validate:"omitempty,min=1,max=100" example:"10"`
}

// PaymentIDResponse represents the response payload containing only a payment ID
type PaymentIDResponse struct {
	ID int `json:"id" example:"1"`
}

// VerifyPaymentRequest represents the request payload for verifying a payment
type VerifyPaymentRequest struct {
	Status string `json:"status" validate:"required,oneof=verified rejected" example:"verified"`
}

// VerifyPaymentResponse represents the response payload for verifying a payment
type VerifyPaymentResponse struct {
	ID      int    `json:"id" example:"1"`
	Status  string `json:"status" example:"verified"`
	Message string `json:"message" example:"Payment verified successfully"`
}