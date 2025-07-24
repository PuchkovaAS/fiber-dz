package main

import (
	"fiber-dz/config"
	"fiber-dz/internal/auth"
	"fiber-dz/internal/pages"
	"fiber-dz/internal/users"
	"fiber-dz/pkg/database"
	"fiber-dz/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()

	logger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(slogfiber.New(logger))
	app.Use(recover.New())

	app.Static("/public", "./public")

	dbpool := database.NewDbPool(dbConfig, logger)
	defer dbpool.Close()

	// repositories
	userRepository := users.NewUsersRepository(
		dbpool,
		logger,
	)

	// services
	authService := auth.NewAuthService(userRepository)

	// handlers
	pages.NewHandler(app, logger)
	auth.NewHandler(app, logger, *authService)
	app.Listen(":3000")
}
