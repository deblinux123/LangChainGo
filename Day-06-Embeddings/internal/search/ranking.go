package search

import (
	"fmt"
	"sort"

	"github.com/deblinux123/LangChainGo/Day-06-Embeddings/internal/similarity"
)

type SearchResult struct {
	Text       string
	Similarity float64
}

func RankDocuments(
	queryVector []float32,
	documents []string,
	documentVector [][]float32,
) ([]SearchResult, error) {
	if len(documents) != len(documentVector) {
		return nil, fmt.Errorf(
			"documents and vectors must have the same length",
		)
	}

	results := make([]SearchResult, 0, len(documents))

	for i, vector := range documentVector {
		score, err := similarity.CosineSimilarity(
			queryVector,
			vector,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"calculate similarity for document %d: %w",
				i,
				err,
			)
		}

		results = append(results, SearchResult{
			Text:       documents[i],
			Similarity: score,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Similarity > results[j].Similarity
	})

	return results, nil
}
