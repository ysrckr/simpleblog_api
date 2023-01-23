package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ysrckr/simpleblog_api/handlers"
)

func ServeRoutes(app *fiber.App) {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	PORT := os.Getenv("DEV_PORT")
	// Route groups
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")

	// Routes
	v1.Get("/healthcheck", handlers.HealthCheck)
	auth.Post("/", handlers.Register)

	app.Listen(PORT)
}



