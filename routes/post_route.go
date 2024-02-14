package routes

import (
	postcontroller "galih/belajar-fiber/controllers/post_controller"
	"galih/belajar-fiber/utils"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App) {
	r := app.Group("/api/v1/posts")

	// JWT Middleware
    r.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: utils.JWTMiddleware,
    }))

	r.Get("/", postcontroller.GetAllPost)
	r.Post("/", postcontroller.CreatePost)
}