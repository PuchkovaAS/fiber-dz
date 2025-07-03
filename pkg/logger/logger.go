package logger

import (
	"fiber-dz/config"
	"log/slog"
	"os"
)

func NewLogger(config *config.LogConfig) *slog.Logger {
	slog.SetLogLoggerLevel(slog.Level(config.Level))

	var logger *slog.Logger

	if config.Format == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return logger
}
