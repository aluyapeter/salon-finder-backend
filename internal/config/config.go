package config

import (
	// "log"
	"os"

	"github.com/aluyapeter/salon-finder-backend/internal/logger"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dburl := os.Getenv("DB_URL")
	if dburl == "" {
		logger.New().Fatal("DB URL is required but not set")
	}

	return Config{
		Port: port,
		DBUrl: dburl,
	}
}