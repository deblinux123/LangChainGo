package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func main() {
	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// create prompt template
	prompt := prompts.NewPromptTemplate(
		`you are a Go programming tutor.
			Explain the folloing topic:
			{{.topic}}
			Keep your explanation simple and provide one example.`,
		[]string{"topic"},
	)

	// create chain
	chain := chains.NewLLMChain(model, prompt)

	// define input
	inputs := map[string]any{
		"topic": "go interface",
	}

	// execute chain
	result, err := chains.Call(
		context.Background(),
		chain,
		inputs,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("AI Response:")
	fmt.Println(result["text"])

}
