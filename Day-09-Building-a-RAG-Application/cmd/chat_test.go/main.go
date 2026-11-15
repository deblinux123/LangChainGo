package main

import (
	"context"
	"log"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/domain"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/service"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/database"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/monitoring"
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

	usageSVC := monitoring.NewUsageService(db)

	chatSVC, err := service.NewChatService(
		"qwen3.5:4b",
		usageSVC,
	)

	if err != nil {
		log.Fatal(err)
	}

	testUser := models.User{
		FullName: "Test User",
		Email:    "test@example.com",
		Role:     models.RoleUser,
		IsActive: true,
	}

	if err := db.Where(
		"email = ?", testUser.Email,
	).FirstOrCreate(&testUser).Error; err != nil {
		log.Fatal("❌ create test user: %w", err)
	}

	response, err := chatSVC.Chat(
		ctx,
		domain.ChatRequest{
			UserID:  testUser.ID,
			Message: "به زبان فارسی و در دو جمله توضیح بده RAG در هوش مصنوعی چیست.",
		},
	)

	if err != nil {
		log.Fatal(
			err,
		)
	}

	log.Println("Model:", response.Model)
	log.Println("Answer:", response.Answer)
	log.Println("Prompt Token:", response.PromptTokens)
	log.Println("Completion Tokens:", response.CompletionTokens)
	log.Println("Total Token:", response.TotalTokens)
}
