package dto

import "time"

// ==================== SHIPMENT DTOs ====================

// CreateShipmentRequest represents the request payload for creating a shipment
type CreateShipmentRequest struct {
	OrderID int    `json:"order_id" validate:"required" example:"1"`
	Kurir   string `json:"kurir" validate:"required,min=1,max=100" example:"JNE"`
	NoResi  string `json:"no_resi" validate:"required,min=1,max=50" example:"JNE1234567890"`
}

// UpdateShipmentRequest represents the request payload for updating a shipment
type UpdateShipmentRequest struct {
	Status string `json:"status" validate:"required,oneof=proses dikirim diterima" example:"dikirim"`
}

// ShipmentResponse represents the response payload for a shipment
type ShipmentResponse struct {
	ID      int       `json:"id" example:"1"`
	OrderID int       `json:"order_id" example:"1"`
	Kurir   string    `json:"kurir" example:"JNE"`
	NoResi  string    `json:"no_resi" example:"JNE1234567890"`
	Status  string    `json:"status" example:"dikirim"`
	Tanggal time.Time `json:"tanggal" example:"2026-04-24T11:00:00Z"`
}

// ShipmentListResponse represents the response payload for a list of shipments
type ShipmentListResponse struct {
	Shipments []ShipmentResponse `json:"shipments"`
}

// ShipmentFilter represents the filter parameters for listing shipments
type ShipmentFilter struct {
	Status  string    `query:"status" validate:"omitempty,oneof=proses dikirim diterima"`
	OrderID *int      `query:"order_id" validate:"omitempty"`
	Kurir   string    `query:"kurir" validate:"omitempty"`
	FromDate time.Time `query:"from_date" validate:"omitempty"`
	ToDate   time.Time `query:"to_date" validate:"omitempty"`
	Page    int       `query:"page" validate:"omitempty,min=1" example:"1"`
	Limit   int       `query:"limit" validate:"omitempty,min=1,max=100" example:"10"`
}

// ShipmentIDResponse represents the response payload containing only a shipment ID
type ShipmentIDResponse struct {
	ID int `json:"id" example:"1"`
}