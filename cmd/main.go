package main

import (
	"fiber-dz/config"
	"fiber-dz/internal/auth"
	"fiber-dz/internal/news"
	"fiber-dz/internal/pages"
	"fiber-dz/internal/users"
	"fiber-dz/pkg/database"
	"fiber-dz/pkg/logger"
	"fiber-dz/pkg/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// repositories
	userRepository := users.NewUsersRepository(
		dbpool,
		logger,
	)
	newsRepository := news.NewNewsRepository(
		dbpool,
		logger,
	)

	// services
	authService := auth.NewAuthService(userRepository)

	// handlers
	pages.NewHandler(app, logger, newsRepository, store)
	auth.NewHandler(app, logger, *authService, store)
	app.Listen(":3000")
}
