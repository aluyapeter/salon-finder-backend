package config

import (
	// "log"
	"os"

	"github.com/aluyapeter/salon-finder-backend/internal/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() Config {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dburl := os.Getenv("DB_URL")
	if dburl == "" {
		logger.New().Fatal("DB URL is required but not implemented")
	}

	return Config{
		Port:  port,
		DBUrl: dburl,
	}
}
