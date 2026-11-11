package main

import (
	"context"
	"log"

	"github.com/deblinux123/LangChainGo/Day-08-Vector-Stores-with-Qdrant/internal/vectorstore"
	"github.com/tmc/langchaingo/schema"
)

func main() {
	ctx := context.Background()

	store, err := vectorstore.NewStore()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] Qdrant vector store initialized successfully.")

	docs := []schema.Document{
		{
			PageContent: "Go is a programming language designed at Google.",
			Metadata: map[string]any{
				"source":   "golang.md",
				"category": "programming",
			},
		},
		{
			PageContent: "Python is widely used in machine learning and artificial intelligence.",
			Metadata: map[string]any{
				"source":   "python.md",
				"category": "ai",
			},
		},
		{
			PageContent: "Golang supports concurrency through goroutines and channels.",
			Metadata: map[string]any{
				"source":   "concurrency.md",
				"category": "programming",
			},
		},
	}

	ids, err := store.AddDocuments(
		ctx,
		docs,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[+] Successfully indexed %d documents.", len(ids))

	for i, id := range ids {
		log.Printf("Document %d ID: %s", i+1, id)
	}
}
