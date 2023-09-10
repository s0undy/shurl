package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/s0undy/shurl/database"
	"github.com/s0undy/shurl/routes"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortURL)
}

func main() {
	log.Println("Loading .env")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connecting to DB")
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	app := fiber.New()

	app.Use(logger.New())

	SetupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
