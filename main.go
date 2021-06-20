package main

import (
	"gotodo/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", Home)

	err := database.Connect()
	if err != nil {
		panic("Connection Database Error!")
	}

	app.Listen(":4000")
}

func Home(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "Hello function",
	})

}
