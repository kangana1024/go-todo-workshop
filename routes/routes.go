package routes

import (
	"gotodo/controllers"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/", Home)
	app.Post("/create", controllers.Create)
	app.Get("/lists", controllers.Lists)
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello ✨")
}
