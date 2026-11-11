package vectorstore

import (
	"net/url"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func NewStore() (vectorstores.VectorStore, error) {
	model, err := ollama.New(
		ollama.WithModel("embeddinggemma:300m"),
	)

	if err != nil {
		return nil, err
	}

	embedder, err := embeddings.NewEmbedder(model)

	if err != nil {
		return nil, err
	}

	qdrantURL, err := url.Parse("http://localhost:6333")

	store, err := qdrant.New(
		qdrant.WithURL(*qdrantURL),
		qdrant.WithCollectionName("langchaingo_day08"),
		qdrant.WithEmbedder(embedder),
	)

	if err != nil {
		return nil, err
	}

	return store, nil
}
