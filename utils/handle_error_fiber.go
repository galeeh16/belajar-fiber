package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var HandleError = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	err = c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
		"code":    code,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
		})
	}

	// Return from handler
	return nil
}
