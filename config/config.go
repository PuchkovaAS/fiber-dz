package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

type DatabaseConfig struct {
	Url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: GetEnv("DATABASE_URL", "NOpe"),
	}
}

func tryConvert[T any](value string) (T, bool) {
	var zero T

	switch any(zero).(type) {
	case string:
		return any(value).(T), true
	case int:
		if intVal, err := strconv.Atoi(value); err == nil {
			return any(intVal).(T), true
		}
	case bool:
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return any(boolVal).(T), true
		}
	case time.Duration:
		if durVal, err := time.ParseDuration(value); err == nil {
			return any(durVal).(T), true
		}
	}
	if str, ok := any(value).(T); ok {
		return str, false
	}
	return zero, false
}

func GetEnv[T any](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	if result, ok := tryConvert[T](value); ok {
		return result
	}
	return defaultValue
}
