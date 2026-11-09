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

	for i, doc := range docs {
		log.Printf("[%d] Document: %v", i+1, doc.Metadata["filename"])
		log.Println("[+] Content:")
		log.Println(doc.PageContent)
	}
}
