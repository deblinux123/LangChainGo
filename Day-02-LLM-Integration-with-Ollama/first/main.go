package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	model, err := ollama.New(ollama.WithModel("qwen3.5:4b"))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		model,
		"tell me about ai in persian",
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	_ = completion
}
