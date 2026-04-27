package dto

import "time"

// ==================== CUSTOMER DTOs ====================

// CustomerRequest represents the request payload for creating/updating a customer
type CustomerRequest struct {
	KodeCust string `json:"kode_cust" validate:"required" example:"CUST-001"`
	Nama     string `json:"nama" validate:"required,min=1,max=200" example:"John Smith"`
	Email    string `json:"email" validate:"required,email" example:"john.smith@example.com"`
	Telepon  string `json:"telepon" validate:"omitempty" example:"+6281234567890"`
	Alamat   string `json:"alamat" validate:"omitempty" example:"Jl. Merdeka No. 123, Jakarta"`
}

// CustomerResponse represents the response payload for a customer
type CustomerResponse struct {
	ID        int       `json:"id" example:"1"`
	KodeCust  string    `json:"kode_cust" example:"CUST-001"`
	Nama      string    `json:"nama" example:"John Smith"`
	Email     string    `json:"email" example:"john.smith@example.com"`
	Telepon   string    `json:"telepon" example:"+6281234567890"`
	Alamat    string    `json:"alamat" example:"Jl. Merdeka No. 123, Jakarta"`
	CreatedAt time.Time `json:"created_at" example:"2026-04-24T10:00:00Z"`
}

// CustomerListResponse represents the response payload for a list of customers
type CustomerListResponse struct {
	Customers []CustomerResponse `json:"customers"`
}

// CustomerFilter represents the filter parameters for listing customers
type CustomerFilter struct {
	Search string `query:"search" validate:"omitempty"`
	Page   int    `query:"page" validate:"omitempty,min=1" example:"1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100" example:"10"`
}

// CustomerIDResponse represents the response payload containing only a customer ID
type CustomerIDResponse struct {
	ID int `json:"id" example:"1"`
}