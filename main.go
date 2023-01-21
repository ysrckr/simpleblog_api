package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysrckr/simpleblog_api/handlers"
)

const PORT = ":4000"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Simple Blog Api Version 1.0.0",
	})

	// Route groups
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")

	// Routes
	v1.Get("/healthcheck", handlers.HealthCheck)
	users.Post("/", handlers.CreateUser)

	app.Listen(PORT)
}
