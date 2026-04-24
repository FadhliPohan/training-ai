package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/response"
)

// JWTClaims represents the custom claims stored inside the JWT token.
type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// contextKey is the fiber Locals key for storing claims.
const contextKey = "claims"

// AuthRequired validates the JWT from the Authorization header (Bearer) or httpOnly cookie.
// On success, the parsed JWTClaims are stored in c.Locals("claims") for downstream handlers.
func AuthRequired(c *fiber.Ctx) error {
	token := extractToken(c)
	if token == "" {
		return response.Unauthorized(c, "")
	}

	claims, err := parseToken(token)
	if err != nil {
		return response.Unauthorized(c, "")
	}

	c.Locals(contextKey, claims)
	return c.Next()
}

// RoleGuard returns a middleware that restricts access to users with at least one of the allowed roles.
func RoleGuard(allowedRoles ...string) fiber.Handler {
	allowed := make(map[string]bool, len(allowedRoles))
	for _, r := range allowedRoles {
		allowed[r] = true
	}

	return func(c *fiber.Ctx) error {
		claims, ok := GetClaims(c)
		if !ok {
			return response.Unauthorized(c, "")
		}
		if !allowed[claims.Role] {
			return response.Forbidden(c, "")
		}
		return c.Next()
	}
}

// ViewerReadOnly blocks write operations (POST, PUT, PATCH, DELETE) for the viewer role.
// Viewers can only perform GET requests.
func ViewerReadOnly(c *fiber.Ctx) error {
	claims, ok := GetClaims(c)
	if !ok {
		return c.Next()
	}
	if claims.Role == "viewer" && c.Method() != fiber.MethodGet {
		return response.Forbidden(c, "Role viewer hanya memiliki akses baca. Tidak bisa melakukan perubahan data.")
	}
	return c.Next()
}

// GetClaims retrieves the parsed JWTClaims from fiber context locals.
func GetClaims(c *fiber.Ctx) (*JWTClaims, bool) {
	claims, ok := c.Locals(contextKey).(*JWTClaims)
	return claims, ok
}

// extractToken retrieves the JWT string from the Authorization header or cookie.
func extractToken(c *fiber.Ctx) string {
	// Try Authorization header first: "Bearer <token>"
	header := c.Get("Authorization")
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	// Fall back to httpOnly cookie
	return c.Cookies("access_token")
}

// parseToken validates the JWT signature and expiry then returns the claims.
func parseToken(tokenStr string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.App.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
