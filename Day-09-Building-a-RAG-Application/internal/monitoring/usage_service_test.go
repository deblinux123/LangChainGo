package monitoring

import (
	"testing"

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
