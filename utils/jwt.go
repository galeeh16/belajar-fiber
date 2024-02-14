package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// middleware
var JWTMiddleware = func(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"message": "Invalid token",
	})
}

type JwtClaim struct {
	Sub string 
	Exp time.Time 
}

// decode token and return data
func DecodeJWT(string) ( error) {
	// jwt := 
	return nil
}