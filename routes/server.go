package routes

import "github.com/gofiber/fiber/v2"

func CreateServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Simple Blog Api Version 1.0.0",
	})

	return app
}
