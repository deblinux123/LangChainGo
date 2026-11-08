package main

import (
	"context"
	"log"

	"github.com/deblinux123/LangChainGo/Day-06-Embeddings/internal/search"
	"github.com/deblinux123/LangChainGo/Day-06-Embeddings/internal/similarity"
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

	vectors, err := embeder.EmbedDocuments(
		ctx,
		texts,
	)

	if err != nil {
		log.Fatal(err)
	}

	score1, err := similarity.CosineSimilarity(vectors[0], vectors[1])

	if err != nil {
		log.Fatal(err)
	}

	score2, err := similarity.CosineSimilarity(vectors[0], vectors[2])

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Similarity 1: %.4f\n", score1)
	log.Printf("Similarity 2: %.4f\n", score2)

	user := "i want to write code for backend"

	userVector, err := embeder.EmbedQuery(ctx, user)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("User Input:", user)
	for i := 0; i < len(texts); i++ {
		score, err := similarity.CosineSimilarity(userVector, vectors[i])

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Document: %s | Similarity: %.4f",
			texts[i], score)
	}

	results, err := search.RankDocuments(
		userVector,
		texts,
		vectors,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ranked search results:")

	for i, result := range results {
		log.Printf(
			"Rank: %d | Similarity: %.4f | Document: %s",
			i+1,
			result.Similarity,
			result.Text,
		)
	}

	topResults := search.TopK(results, 2)

	log.Println("Top-K Results:")

	for i, result := range topResults {
		log.Printf(
			"Rank: %d | Score: %.4f | Document: %s",
			i+1,
			result.Similarity,
			result.Text,
		)
	}

}
