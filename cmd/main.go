package main

import (
	"fiber-dz/config"
	"fiber-dz/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()

	app := fiber.New()

	app.Use(recover.New())

	pages.NewHandler(app)
	app.Listen(":3000")
}
