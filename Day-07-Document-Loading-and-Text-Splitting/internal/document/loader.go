package document

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadeFile(path string) (Document, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return Document{}, fmt.Errorf("read file: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(path))

	if ext != ".txt" && ext != ".md" {
		return Document{}, fmt.Errorf("unsupported file format: %s", ext)
	}

	doc := Document{
		PageContent: string(content),
		Metadata: map[string]any{
			"source":   path,
			"filename": filepath.Base(path),
			"type":     strings.TrimPrefix(ext, "."),
		},
	}

	return doc, nil
}
