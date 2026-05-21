package api

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL string
	Timeout time.Duration
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("load .env file: %w", err)
	}

	timeoutSeconds, err := getEnvAsInt("API_TIMEOUT_SECONDS", 10)
	if err != nil {
		return Config{}, err
	}

	return Config{
		BaseURL: getEnv("API_BASE_URL", "https://api.open-meteo.com"),
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}, nil
}


func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
} 


func getEnvAsInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("parse env %s as int: %w", key, err)
	}

	return parsed, nil
}