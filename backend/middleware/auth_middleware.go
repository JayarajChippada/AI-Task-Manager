package middleware

import (
	"net/http"

	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from cookies instead of headers
	token := c.Cookies("token") // Fetch token from cookie named "token"
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authentication token"})
	}

	// Verify JWT token
	userID, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Store userID in locals for use in controllers
	c.Locals("userID", userID)

	return c.Next()
}
