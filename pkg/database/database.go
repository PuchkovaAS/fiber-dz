package database

import (
	"context"
	"fiber-dz/config"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDbPool(
	config *config.DatabaseConfig,
	logger *slog.Logger,
) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error("Не удалось подключиться к БД")
		panic(err)
	}
	logger.Info("Подключились к БД")
	return dbpool
}
