package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
		ollama.WithThink(true),
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Create model:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Local LLM CLI")
	fmt.Println("Type 'exit' to quit.")

	for {
		fmt.Print("\nYou: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Read input:", err)
			return
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			return
		}

		if input == "" {
			continue
		}

		ctx, cancel := context.WithTimeout(
			context.Background(),
			60*time.Second,
		)

		response, err := model.GenerateContent(
			ctx,
			[]llms.MessageContent{
				llms.TextParts(
					llms.ChatMessageTypeHuman,
					input,
				),
			},

			llms.WithTemperature(0.6),
			llms.WithMaxTokens(1500),
			llms.WithTopP(0.9),
			llms.WithReturnThinking(true),
		)
		defer cancel()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Generate response:", err)
			continue
		}

		fmt.Printf("DEBUG: %#v\n", response.Choices[0].GenerationInfo)

		fmt.Print("Assistant: ")

		for i, choice := range response.Choices {
			fmt.Printf("\nChoice %d\n", i)
			fmt.Printf("Content: %q\n", choice.Content)
			fmt.Printf("ReasoningContent: %q\n", choice.ReasoningContent)
			fmt.Printf("GenerationInfo: %#v\n", choice.GenerationInfo)
		}
	}
}
