package main

import (
	"context"
	"log"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/domain"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/service"
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

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ DataBase migration complited.")

	if err := database.SeedPlan(db); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Initial plans created successfully.")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Minute,
	)

	defer cancel()

	chatSVC, err := service.NewModelService("qwen3.5:4b")

	if err != nil {
		log.Fatal(err)
	}

	response, err := chatSVC.Chat(ctx, domain.ChatRequest{
		Message: "به زبان فارسی و در دو جمله توضیح بده RAG در هوش مصنوعی چیست.",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Model:", response.Model)
	log.Println("Answer:", response.Answer)
	log.Println("Prompt Token:", response.PromptTokens)
	log.Println("Completion Tokens:", response.CompletionTokens)
	log.Println("Total Token:", response.TotalTokens)
}
