package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func main() {
	template := prompts.NewPromptTemplate(
		`Classify the sentiment of the input.

Use only one of these labels:
Positive
Negative
Neutral

Examples:

Input: I love this product.
Output: Positive

Input: This service is terrible.
Output: Negative

Input: The package arrived yesterday.
Output: Neutral

Now classify this input:

Input: {{.text}}
Output:`,
		[]string{"text"},
	)

	input := "The quality is excellent!."

	prompt, err := template.Format(map[string]any{
		"text": input,
	})

	if err != nil {
		log.Fatal(err)
	}

	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		50*time.Second,
	)

	defer cancel()

	response, err := model.GenerateContent(
		ctx,
		[]llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeHuman, prompt),
		},
		llms.WithTemperature(0.7),
	)

	if err != nil {
		log.Fatal(err)
	}

	if len(response.Choices) == 0 {
		log.Fatal(err)
	}

	fmt.Println("Input:", input)
	fmt.Println("Sentiment:", strings.TrimSpace(response.Choices[0].Content))
}
