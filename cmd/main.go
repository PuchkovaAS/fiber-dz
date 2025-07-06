package main

import (
	"fiber-dz/config"
	"fiber-dz/internal/pages"
	"fiber-dz/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	logger := logger.NewLogger(logConfig)

	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]any{
		"AddHash": func(c string) string {
			return "#" + c
		},
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(slogfiber.New(logger))
	app.Use(recover.New())

	pages.NewHandler(app, logger)
	app.Listen(":3000")
}
