package document

import (
	"unicode/utf8"

	"github.com/tmc/langchaingo/textsplitter"
)

func NewSplitter() textsplitter.RecursiveCharacter {
	return textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(700),
		textsplitter.WithChunkOverlap(100),
		textsplitter.WithLenFunc(func(text string) int {
			return utf8.RuneCountInString(text)
		}),
		textsplitter.WithKeepSeparator(false),
		textsplitter.WithSeparators([]string{
			"\n\n",
			"\n",
			"۔",
			".",
			"؟",
			"?",
			"!",
			"؛",
			"،",
			" ",
			"",
		}),
	)
}
