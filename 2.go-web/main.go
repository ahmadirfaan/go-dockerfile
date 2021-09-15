package main

import (
	"github.com/gofiber/fiber/v2"
)
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hai gaes ini golang fiber")
	})

	app.Listen(":8000")
}