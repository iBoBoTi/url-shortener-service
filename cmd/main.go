package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.SShortenURL)
}

func main() {
	err := godotenv.Load()
	app := fiber.New()
}
