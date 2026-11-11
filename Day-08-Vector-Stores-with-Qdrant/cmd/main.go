package main

import (
	"context"
	"log"

	"github.com/deblinux123/LangChainGo/Day-08-Vector-Stores-with-Qdrant/internal/vectorstore"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
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

	query := "How does go handle concurrent tasks?"
	// results, err := store.SimilaritySearch(
	// 	ctx,
	// 	query,
	// 	2,
	// 	vectorstores.WithFilters(map[string]any{
	// 		"must": map[string]any{
	// 			"key": "category",
	// 			"match": map[string]any{
	// 				"value": "programming",
	// 			},
	// 		},
	// 	}),
	// )

	retriver := vectorstores.ToRetriever(
		store,
		2,
	)

	results, err := retriver.GetRelevantDocuments(
		ctx,
		query,
	)

	if err != nil {
		log.Fatal(err)
	}

	// log.Println("[+] Search Results:")
	log.Println("[+] Retriever Results:")

	for i, doc := range results {
		log.Printf("\n--- Result %d ---", i+1)
		log.Println("Content:", doc.PageContent)
		log.Println("Metadata:", doc.Metadata)
	}
}
