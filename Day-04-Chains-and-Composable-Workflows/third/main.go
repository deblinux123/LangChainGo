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
	ctx, cancle := context.WithTimeout(
		context.Background(),
		120*time.Second,
	)

	defer cancle()

	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatalf("create model: %w", err)
	}

	summeryPrompt := prompts.NewPromptTemplate(
		`Summerizer the following text in maximum 3 sentences
		
		Text:
		{{.text}}
		`,
		[]string{"text"},
	)

	summeryChain := chains.NewLLMChain(
		model,
		summeryPrompt,
	)

	summeryChain.OutputKey = "summery"

	categoryPrompt := prompts.NewPromptTemplate(
		`Classify the following summery exactly one category:
		AI, Prgramming, Science, Business, or Other.
		
		Return only the category name.
		
		Summer:
		{{.summery}}`,
		[]string{"summery"},
	)

	categoryChain := chains.NewLLMChain(
		model,
		categoryPrompt,
	)

	categoryChain.OutputKey = "category"

	// inputText := `
	// Go is a statically typed programming language
	// developed at Google. It supports concurrency
	// through goroutines and channels. Go is widely
	// used for backend services and cloud infrastructure.
	// `

	inputText := `
Retrieval-Augmented Generation (RAG) combines
information retrieval with language models.
It uses vector databases to retrieve relevant
documents and provide context to an LLM.
`
	summerResult, err := chains.Call(
		ctx,
		summeryChain,
		map[string]any{
			"text": inputText,
		},
	)

	if err != nil {
		log.Fatalf("Summerization failed: %v", err)
	}

	summery, ok := summerResult["summery"].(string)

	if !ok {
		log.Fatal("Invalid summry output.")
	}

	categoryResult, err := chains.Call(
		ctx,
		categoryChain,
		map[string]any{
			"summery": summery,
		},
	)

	if err != nil {
		log.Fatalf("categorization failed: %v", err)
	}

	category, ok := categoryResult["category"].(string)

	if !ok {
		log.Fatal("Invalid category output")
	}

	fmt.Println("========== WORKFLOW RESULT ==========")
	fmt.Println("Summary:", summery)
	fmt.Println("Category:", category)
}
