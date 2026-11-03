package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "flashhelper"),
		DBPassword: getEnv("DB_PASSWORD", "flashhelper-dev-password"),
		DBName:     getEnv("DB_NAME", "flashhelper"),
		JWTSecret:  getEnv("JWT_SECRET_KEY", "flashhelper-local-development-secret-change-before-deploying"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists && value != "" {
		return value
	}

	return fallback
}
