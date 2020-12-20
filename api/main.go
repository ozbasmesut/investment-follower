package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", investorHandler)

	app.Listen(":6161")

}

func investorHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, Investors 💸!")
}
