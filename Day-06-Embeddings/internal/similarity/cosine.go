package similarity

import (
	"errors"
	"math"
)

func CosineSimilarity(a, b []float32) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("Vectors must have the same dimensions")
	}

	if len(a) == 0 {
		return 0, errors.New("Vector cannot be empty")
	}

	var dotProduct float64
	var magnitudeA float64
	var magnitudeB float64

	for i := range a {
		x := float64(a[i])
		y := float64(b[i])

		dotProduct += x * y
		magnitudeA += x * x
		magnitudeB += y * y
	}

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0, errors.New("cannot calculate similarity for zero vectors.")
	}

	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB)), nil
}
