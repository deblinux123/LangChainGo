package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	messages := []llms.MessageContent{
		llms.TextParts(
			llms.ChatMessageTypeSystem,
			"You are a Go programming tutor. Explain concepts clearly and use short examples.",
		),
		llms.TextParts(
			llms.ChatMessageTypeHuman,
			"What is a Go (golang) interface",
		),
	}

	response, err := model.GenerateContent(
		ctx,
		messages,
	)

	if err != nil {
		log.Fatal(err)
	}

	if len(response.Choices) == 0 {
		log.Fatalf("No response generated.")
	}

	fmt.Printf("AI:\n%s", response.Choices[0].Content)
}
