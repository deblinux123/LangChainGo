package main

import (
	"log"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/database"
)

func main() {
	db, err := database.NewSqlite()

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	log.Println("✅ Database is ready.")
}
