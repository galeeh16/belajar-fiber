package routes

import (
	contactcontroller "galih/belajar-fiber/controllers/contact_controller"
	"galih/belajar-fiber/utils"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func ContactRoute(app *fiber.App) {
	r := app.Group("/api/v1/contacts")

	// JWT Middleware
    r.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: utils.JWTMiddleware,
    }))

	r.Get("/", contactcontroller.GetAllContact)
	r.Post("/", contactcontroller.CreateContact)
}