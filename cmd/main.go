package main

import (
	"fiber-dz/config"
	"fiber-dz/internal/pages"
	"fiber-dz/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	logger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(slogfiber.New(logger))
	app.Use(recover.New())

	app.Static("/public", "./public")

	pages.NewHandler(app, logger)
	app.Listen(":3000")
}
