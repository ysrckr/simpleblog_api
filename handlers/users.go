package handlers

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}
