package subscription

import (
	"fmt"
	"testing"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dsn := fmt.Sprintf(
		"file:%s?mode=memory&cache=shared",
		t.Name(),
	)

	db, err := gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to get sql database: %v", err)
	}

	t.Cleanup(func() {
		if err := sqlDB.Close(); err != nil {
			t.Errorf("failed to close test database: %v", err)
		}
	})

	err = db.AutoMigrate(
		&models.User{},
		&models.Plan{},
		&models.Subscription{},
		&models.Usage{},
	)
	if err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}

func TestCreateSubscription(t *testing.T) {
	db := setupTestDB(t)

	// Create a test user.
	user := models.User{
		FullName:     "Test User",
		Email:        "test@example.com",
		PasswordHash: "test-hash",
		Role:         models.RoleUser,
		IsActive:     true,
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	// Create the Free plan.
	plan := models.Plan{
		Code:                models.PlanFree,
		Name:                "Free",
		DailyMessageLimit:   20,
		MonthlyTokenLimit:   100000,
		CanAccessAllLessons: false,
		CanAccessAdvanced:   false,
		IsActive:            true,
	}

	if err := db.Create(&plan).Error; err != nil {
		t.Fatalf("failed to create test plan: %v", err)
	}

	service := NewService(db)

	// Create a Free subscription.
	subscription, err := service.CreateSubscription(
		user.ID,
		models.PlanFree,
	)
	if err != nil {
		t.Fatalf("failed to create subscription: %v", err)
	}

	if subscription.UserID != user.ID {
		t.Errorf("expected user ID %d, got %d", user.ID, subscription.UserID)
	}

	if subscription.PlanID != plan.ID {
		t.Errorf("expected plan ID %d, got %d", plan.ID, subscription.PlanID)
	}

	if subscription.Status != models.SubscriptionActive {
		t.Errorf("expected active status, got %s", subscription.Status)
	}

	if subscription.EndsAt != nil {
		t.Error("expected Free subscription to have no expiration date")
	}

	// A second active subscription must be rejected.
	_, err = service.CreateSubscription(user.ID, models.PlanFree)
	if err == nil {
		t.Error("expected duplicate active subscription to fail")
	}
}
