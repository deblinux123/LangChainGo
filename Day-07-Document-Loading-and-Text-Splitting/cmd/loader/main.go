package main

import (
	"log"

	"github.com/deblinux123/LangChainGo/Day-07-Document-Loading-and-Text-Splitting/internal/document"
)

func main() {
	docs, err := document.LoadDirectory("/home/toor/Desktop/all/LangChainGo/Day-07-Document-Loading-and-Text-Splitting/data")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] File are loaded successfully.")
	log.Printf("[+] %d Documents are loaded.", len(docs))

	splitter := document.NewSplitter()

	validDocs := document.ValidateDocuments(docs)

	for i, doc := range docs {
		log.Printf("[%d] Document: %v", i+1, doc.Metadata["filename"])
		chunks, err := splitter.SplitText(doc.PageContent)

		if err != nil {
			log.Fatal(err)
		}

		for i, chunk := range chunks {
			log.Printf("\n--- Chunk %d ---\n", i+1)
			log.Println(chunk)
		}
	}

	log.Printf("[+] %d Valid Documents.", len(validDocs))

}
