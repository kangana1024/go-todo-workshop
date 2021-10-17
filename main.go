package main

import (
	"gotodo/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"gotodo/routes"
)

func main() {
	if os.Getenv("CHECK_ENV") != "1" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error : %+v!", err)
		}
	}
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		panic("Connection Database Error!")
	}

	routes.Init(app)

	app.Listen(":4281")
}
