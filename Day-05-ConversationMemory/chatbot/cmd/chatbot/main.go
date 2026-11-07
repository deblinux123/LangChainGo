package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deblinux123/LangChainGo/Day-05-ConversationMemory/chatbot/internal/chat"
	"github.com/deblinux123/LangChainGo/Day-05-ConversationMemory/chatbot/internal/config"
	"github.com/deblinux123/LangChainGo/Day-05-ConversationMemory/chatbot/internal/redisclient"
	"github.com/tmc/langchaingo/llms/ollama"
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

	model, err := ollama.New(
		ollama.WithModel(cfg.OllamaModel),
	)

	if err != nil {
		log.Fatal(err)
	}

	chatService := chat.NewService(model, redisClient)

	ctx := context.Background()

	scanner := bufio.NewScanner(os.Stdin)

	sessionID := "farid-session"

	fmt.Println("Chat Started. Type 'exit' to quit.")

	for {
		fmt.Print("\nYou: ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			break
		}

		answer, err := chatService.Chat(ctx, sessionID, input)

		if err != nil {
			log.Println("Chat error: %w", err)
			continue
		}

		fmt.Println("AI: ", answer)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Input error: %v", err)
	}
}
