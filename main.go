package main

import (
	"fmt"
	"galih/belajar-fiber/routes"
	"galih/belajar-fiber/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	// load .env file
	utils.LoadEnv()
	// init logrus logger
	utils.InitLogger()
	// connect db
	utils.ConnectDB()

	// init fiber app
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
		ErrorHandler: utils.HandleError,
	})

	// use swagger docs
	// app.Use(swagger.New(swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: "./docs/swagger.json",
	// 	Path:     "swagger",
	// 	Title:    "Swagger API Docs",
	// }))

	// enabling cors
	app.Use(cors.New())

	// enable rate limiter
	app.Use(limiter.New(limiter.Config{
		Max: 60, // 60 kali dalam 60 detik (bawah)
		Expiration: time.Second * 60,
	}))

	// // Logging remote IP and Port
	// app.Use(logger.New(logger.Config{
	// 	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	// 	TimeZone: "Asia/Jakarta",
	// }))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Golang fiber v2 auth"})
	})

	// register routes
	routes.UserRoute(app)
	routes.ContactRoute(app)
	routes.PostRoute(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
