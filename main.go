package main

import (
	"gotodo/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Use(cors.New(cors.ConfigDefault))
	err := database.Connect()
	if err != nil {
		panic("Connection Database Error!")
	}

	routes.Init(app)

	app.Listen("0.0.0.0:4281")
}
