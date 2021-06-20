package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", Home)

	app.Listen(":4000")
}

func Home(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "Hello function",
	})

}
