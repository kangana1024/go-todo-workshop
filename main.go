package main

import (
	"gotodo/database"

	"github.com/gofiber/fiber/v2"

	"gotodo/routes"
)

func main() {
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		panic("Connection Database Error!")
	}

	routes.Init(app)

	app.Listen(":4000")
}
