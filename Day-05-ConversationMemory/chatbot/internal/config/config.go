package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	OllamaURL   string
	OllamaModel string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return &Config{}, fmt.Errorf("Load Env file: %w", err)
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))

	if err != nil {
		return &Config{}, fmt.Errorf("Invalid Redis DB: %w", err)
	}

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),

		RedisAddr:     getEnv("REDIS_ADDR", "127.0.0.1:6379"),
		RedisPassword: os.Getenv("REDIS_PASS"),
		RedisDB:       redisDB,
		OllamaURL:     getEnv("OLLAMA_URL", "http://localhost:11434"),
		OllamaModel:   getEnv("OLLAMA_MODEL", "gemma4:e2b"),
	}

	if cfg.RedisAddr == "" {
		return &Config{}, fmt.Errorf("Redis address is required.")
	}

	if cfg.OllamaURL == "" {
		return &Config{}, fmt.Errorf("Ollama url is required")
	}

	if cfg.OllamaModel == "" {
		return &Config{}, fmt.Errorf("Ollama model is required.")
	}

	return cfg, nil

}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
