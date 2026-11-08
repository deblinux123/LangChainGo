package main

import (
	"context"
	"log"
	"time"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	ctx := context.Background()

	model, err := ollama.New(
		ollama.WithModel("embeddinggemma:300m"),
	)

	if err != nil {
		log.Fatal(err)
	}

	embeder, err := embeddings.NewEmbedder(model)

	if err != nil {
		log.Fatal(err)
	}

	texts := []string{
		"I love programming with Go.",
		"I enjoy writing code in Golang.",
		"The weather is very cold today.",
	}

	for i, text := range texts {
		start := time.Now()

		vector, err := embeder.EmbedQuery(ctx, text)

		if err != nil {
			log.Fatal(err)
		}

		elapsed := time.Since(start)

		log.Printf(
			"Document %d | Dimensions: %d | Time: %s\n",
			i+1,
			len(vector),
			elapsed,
		)
	}
}
