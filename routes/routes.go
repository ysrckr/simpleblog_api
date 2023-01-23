package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysrckr/simpleblog_api/handlers"
)

func ServeRoutes(app *fiber.App, port string) {
	// Route groups
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")

	// Routes
	v1.Get("/healthcheck", handlers.HealthCheck)
	auth.Post("/", handlers.Register)

	app.Listen(port)
}



