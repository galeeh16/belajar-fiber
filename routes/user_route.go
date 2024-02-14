package routes

import (
	authcontroller "galih/belajar-fiber/controllers/auth_controller"
	usercontroller "galih/belajar-fiber/controllers/user_controller"
	"galih/belajar-fiber/utils"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// no middleware
	app.Post("/api/v1/users", usercontroller.CreateUser)
	app.Post("/api/v1/users/login", authcontroller.Login)

	r := app.Group("/api/v1/users")

	// JWT Middleware
    r.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: utils.JWTMiddleware,
    }))

	r.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON("user route")
	})

	r.Get("/", usercontroller.GetAllUser)
	r.Put("/:id", usercontroller.UpdateUser)
	r.Delete("/:id", usercontroller.DeleteUser)
}