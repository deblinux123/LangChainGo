package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func main() {
	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatalf("Create model:", err)
	}

	// create prompt template
	prompt := prompts.NewPromptTemplate(
		`You are a professional text summerizer

		summerize the following text in persian language in maximum 3 sentence and keep the main ideas of the text.:
		{{.text}}

		requirements:
		- Preserve the main ideas.
		- User clear and simple language.
		- Maximum 3 sentence.
`,
		[]string{"text"},
	)

	summerizationChain := chains.NewLLMChain(
		model,
		prompt,
	)

	// define input
	inputs := map[string]any{
		"text": `The Go programming language, often referred to as Golang, is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and strong support for concurrent programming. Go features a statically typed syntax, garbage collection, and built-in support for web development and networking. It has gained popularity for its performance, ease of use, and suitability for building scalable applications. it has built-in support for testing and profiling, making it a preferred choice for many developers. The language's standard library provides a wide range of packages for various tasks, and its tooling ecosystem is robust, with tools for formatting, linting, and dependency management. Overall, Go is a versatile language that continues to evolve and attract a growing community of developers.`,
	}

	fmt.Println("Input Key:")
	fmt.Println(summerizationChain.GetInputKeys())

	fmt.Println("Output Key:")
	fmt.Println(summerizationChain.GetOutputKeys())
	// execute chain
	ctx, cancel := context.WithTimeout(
		context.Background(),
		60*time.Second,
	)

	defer cancel()

	result, err := chains.Call(
		ctx,
		summerizationChain,
		inputs,
	)

	if err != nil {
		log.Fatalf("Call chain: %v", err)
	}

	fmt.Println("AI Response:")
	fmt.Println(result["text"])

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}

}
