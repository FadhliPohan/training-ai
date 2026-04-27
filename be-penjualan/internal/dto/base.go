package dto

// StandardResponse represents the standard API response structure
type StandardResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// PaginationMeta contains pagination information
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// PaginatedResponse represents a response with pagination
type PaginatedResponse struct {
	StandardResponse
	Meta PaginationMeta `json:"meta"`
}
