package main

import (
	"log"

	"github.com/deblinux123/LangChainGo/Day-07-Document-Loading-and-Text-Splitting/internal/document"
)

func main() {
	doc, err := document.LoadeFile("/home/toor/Desktop/all/LangChainGo/Day-07-Document-Loading-and-Text-Splitting/ai.md")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Document content:")
	log.Println(doc.PageContent)

	log.Println("\nDocument Metadata:")

	for k, v := range doc.Metadata {
		log.Printf("%s: %v\n", k, v)
	}
}
