package response

import "github.com/gofiber/fiber/v2"

// Standard is the unified API response envelope used across all endpoints.
type Standard struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// OK sends a 200 success response.
func OK(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Standard{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created sends a 201 created response.
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Standard{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// BadRequest sends a 400 validation error response.
func BadRequest(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Standard{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// Unauthorized sends a 401 unauthenticated response.
func Unauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Sesi tidak valid atau sudah kedaluwarsa. Silakan login kembali."
	}
	return c.Status(fiber.StatusUnauthorized).JSON(Standard{
		Success: false,
		Message: message,
	})
}

// Forbidden sends a 403 access denied response.
func Forbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Anda tidak memiliki izin untuk mengakses resource ini."
	}
	return c.Status(fiber.StatusForbidden).JSON(Standard{
		Success: false,
		Message: message,
	})
}

// NotFound sends a 404 not found response.
func NotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Data yang diminta tidak ditemukan."
	}
	return c.Status(fiber.StatusNotFound).JSON(Standard{
		Success: false,
		Message: message,
	})
}

// Conflict sends a 409 conflict response.
func Conflict(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusConflict).JSON(Standard{
		Success: false,
		Message: message,
	})
}

// InternalServerError sends a 500 server error response (without leaking internals).
func InternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Standard{
		Success: false,
		Message: "Terjadi kesalahan pada server. Silakan coba beberapa saat lagi.",
	})
}
