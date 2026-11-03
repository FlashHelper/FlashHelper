package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	AuthUsername string
	AuthPassword string
	SessionTTL   time.Duration
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	sessionTTL, err := time.ParseDuration(getEnv("SESSION_TTL", "24h"))
	if err != nil || sessionTTL <= 0 {
		log.Println("Invalid SESSION_TTL, using default 24h")
		sessionTTL = 24 * time.Hour
	}

	return &Config{
		AppPort:      getEnv("APP_PORT", "8080"),
		AuthUsername: getEnv("AUTH_USERNAME", "admin"),
		AuthPassword: getEnv("AUTH_PASSWORD", "flashhelper-dev-password"),
		SessionTTL:   sessionTTL,
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists && value != "" {
		return value
	}

	return fallback
}
