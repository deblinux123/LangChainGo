package document

import "strings"

func ValidateDocuments(docs []Document) []Document {
	var valid []Document

	for _, doc := range docs {
		if strings.TrimSpace(doc.PageContent) == "" {
			continue
		}

		if doc.Metadata == nil {
			continue
		}

		valid = append(valid, doc)
	}

	return valid
}
