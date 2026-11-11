package main

import (
	"log"

	"github.com/deblinux123/LangChainGo/Day-08-Vector-Stores-with-Qdrant/internal/vectorstore"
)

func main() {
	store, err := vectorstore.NewStore()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] Qdrant vector store initialized successfully.")

	_ = store
}
