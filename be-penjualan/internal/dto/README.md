# DTO (Data Transfer Objects)

This directory contains Data Transfer Objects (DTOs) for the InsightFlow API. DTOs are used to:

1. Define the structure of request and response payloads
2. Provide validation rules for incoming data
3. Document API endpoints with examples
4. Separate API contracts from internal domain models

## Structure

Each module has its own DTO file:

- `auth.go` - Authentication related DTOs (login, register, profile)
- `produk.go` - Product related DTOs (create, update, list)
- `customer.go` - Customer related DTOs (create, update, list)
- `user.go` - User management DTOs (create, update, deactivate)
- `order.go` - Order related DTOs (create, confirm, cancel)
- `payment.go` - Payment related DTOs (create, verify)
- `shipment.go` - Shipment related DTOs (create, update)
- `report.go` - Report related DTOs (generate, data structures)
- `settings.go` - Settings related DTOs (Telegram, anomaly config)
- `base.go` - Base DTOs (standard response, pagination)

## Usage

DTOs are used in handlers to:

1. Parse and validate incoming request data
2. Structure outgoing response data
3. Provide type safety and documentation

Example usage in a handler:

```go
func CreateUser(c *fiber.Ctx) error {
    var req dto.UserRequest
    if err := c.BodyParser(&req); err != nil {
        return response.BadRequest(c, "Invalid request format", nil)
    }

    // Validate the request
    if err := validate.Struct(req); err != nil {
        return response.BadRequest(c, "Validation failed", err)
    }

    // Process the request...
    
    resp := dto.UserResponse{
        ID:    userID,
        Nama:  req.Nama,
        Email: req.Email,
        // ...
    }
    
    return response.Created(c, "User created successfully", resp)
}
```

## Validation

All DTOs include validation tags using the `go-playground/validator` library:

- `required` - Field is required
- `min`/`max` - Minimum/maximum values or lengths
- `email` - Valid email format
- `uuid` - Valid UUID format
- `oneof` - Value must be one of the specified options
- `datetime` - Valid datetime format

## Documentation

DTOs include Swagger/OpenAPI documentation tags:

- `example` - Example values for documentation
- Struct and field comments for descriptions

This allows for automatic generation of API documentation.