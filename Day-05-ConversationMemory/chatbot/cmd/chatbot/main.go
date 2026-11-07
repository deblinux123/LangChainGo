package main

import (
	"log"

	"github.com/deblinux123/LangChainGo/Day-05-ConversationMemory/chatbot/internal/config"
	"github.com/deblinux123/LangChainGo/Day-05-ConversationMemory/chatbot/internal/redisclient"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration loaded.")

	redisClient, err := redisclient.NewClient(
		redisclient.Config{
			Addr:     cfg.RedisAddr,
			Password: cfg.RedisPassword,
			DB:       cfg.RedisDB,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	defer redisClient.Close()

	log.Println("Redis connected successfully.")
	log.Println("Ollama Model:", cfg.OllamaModel)
}
