package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	ctx := context.Background()

	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatal(err)
	}

	messages := []llms.MessageContent{
		llms.TextParts(
			llms.ChatMessageTypeSystem,
			"you are a helpful assistant. and always answer in persian language.",
		),
		llms.TextParts(
			llms.ChatMessageTypeHuman,
			"my name is farid and i'm backemd developer with golang",
		),
	}

	response, err := model.GenerateContent(
		ctx,
		messages,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Assistant:", response.Choices[0].Content)

	messages = []llms.MessageContent{
		llms.TextParts(
			llms.ChatMessageTypeSystem,
			"you are a helpful assistant. and always answer in persian language.",
		),

		llms.TextParts(
			llms.ChatMessageTypeHuman,
			"What is my name?",
		),
	}

	response, err = model.GenerateContent(
		ctx,
		messages,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Assistant:", response.Choices[0].Content)
}
