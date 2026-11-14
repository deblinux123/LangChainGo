package monitoring

import (
	"testing"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUsageTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(
		sqlite.Open("file:usage_test?mode=memory&cache=shared"),
		&gorm.Config{},
	)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Usage{},
	); err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}

func TestRecordUsage(t *testing.T) {
	db := setupUsageTestDB(t)

	// Create a test user.
	user := models.User{
		FullName:     "Test User",
		Email:        "usage-test@example.com",
		PasswordHash: "test-hash",
		Role:         models.RoleUser,
		IsActive:     true,
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	service := NewUsageService(db)

	usage := &models.Usage{
		UserID:       user.ID,
		Feature:      "conversation",
		Model:        "qwen3.5:4b",
		InputTokens:  150,
		OutputTokens: 300,
		LatencyMs:    1200,
		Status:       models.UsageSuccess,
	}

	if err := service.RecordUsage(usage); err != nil {
		t.Fatalf("failed to record usage: %v", err)
	}

	if usage.ID == 0 {
		t.Fatal("expected usage ID to be assigned")
	}

	// Verify the saved record.
	var saved models.Usage
	if err := db.First(&saved, usage.ID).Error; err != nil {
		t.Fatalf("failed to retrieve usage: %v", err)
	}

	if saved.UserID != user.ID {
		t.Errorf("expected user ID %d, got %d", user.ID, saved.UserID)
	}

	if saved.InputTokens != 150 {
		t.Errorf("expected 150 input tokens, got %d", saved.InputTokens)
	}

	if saved.OutputTokens != 300 {
		t.Errorf("expected 300 output tokens, got %d", saved.OutputTokens)
	}

	if saved.Status != models.UsageSuccess {
		t.Errorf("expected status %s, got %s", models.UsageSuccess, saved.Status)
	}
}

func TestRecordUsageRejectsNil(t *testing.T) {
	db := setupUsageTestDB(t)
	service := NewUsageService(db)

	if err := service.RecordUsage(nil); err == nil {
		t.Fatal("expected an error when usage is nil")
	}
}

func TestUsageStatistics(t *testing.T) {
	db := setupUsageTestDB(t)

	user := models.User{
		FullName:     "Statistics User",
		Email:        "statistics@example.com",
		PasswordHash: "test-hash",
		Role:         models.RoleUser,
		IsActive:     true,
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	service := NewUsageService(db)

	now := time.Now().UTC()

	startOfDay := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, time.UTC,
	)

	startOfMonth := time.Date(
		now.Year(), now.Month(), 1,
		0, 0, 0, 0, time.UTC,
	)

	usages := []models.Usage{
		{
			UserID:       user.ID,
			Feature:      "conversation",
			Model:        "qwen3.5:4b",
			InputTokens:  100,
			OutputTokens: 50,
			Status:       models.UsageSuccess,
			CreatedAt:    startOfDay.Add(time.Hour),
		},
		{
			UserID:       user.ID,
			Feature:      "reading",
			Model:        "qwen3.5:4b",
			InputTokens:  200,
			OutputTokens: 100,
			Status:       models.UsageSuccess,
			CreatedAt:    startOfDay.Add(2 * time.Hour),
		},
		{
			UserID:       user.ID,
			Feature:      "conversation",
			Model:        "qwen3.5:4b",
			InputTokens:  500,
			OutputTokens: 500,
			Status:       models.UsageFailed,
			CreatedAt:    startOfDay.Add(3 * time.Hour),
		},
		{
			UserID:       user.ID,
			Feature:      "vocabulary",
			Model:        "qwen3.5:4b",
			InputTokens:  10,
			OutputTokens: 20,
			Status:       models.UsageSuccess,
			CreatedAt:    startOfMonth.Add(time.Hour),
		},
	}

	if err := db.Create(&usages).Error; err != nil {
		t.Fatalf("failed to create usage records: %v", err)
	}

	// Test daily message count.
	dailyCount, err := service.GetDailyMessageCount(user.ID)
	if err != nil {
		t.Fatalf("failed to get daily message count: %v", err)
	}

	if dailyCount != 2 {
		t.Errorf("expected 2 daily messages, got %d", dailyCount)
	}

	// Test monthly token usage.
	monthlyTokens, err := service.GetMonthlyTokenUsage(user.ID)
	if err != nil {
		t.Fatalf("failed to get monthly token usage: %v", err)
	}

	// 150 + 300 + 30 = 480
	if monthlyTokens != 480 {
		t.Errorf("expected 480 monthly tokens, got %d", monthlyTokens)
	}
}
