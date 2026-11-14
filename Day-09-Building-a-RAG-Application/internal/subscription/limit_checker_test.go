package subscription

import (
	"testing"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/monitoring"
)

func TestCheckDailyMessageLimit(t *testing.T) {
	db := setupTestDB(t)

	// Create a test user.
	user := models.User{
		FullName:     "Limit Test User",
		Email:        "limit-test@example.com",
		PasswordHash: "test-hash",
		Role:         models.RoleUser,
		IsActive:     true,
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}

	// Create a Free plan with a limit of 2 messages.
	plan := models.Plan{
		Code:              models.PlanFree,
		Name:              "Free",
		DailyMessageLimit: 2,
		MonthlyTokenLimit: 100000,
		IsActive:          true,
	}

	if err := db.Create(&plan).Error; err != nil {
		t.Fatalf("create plan: %v", err)
	}

	// Create an active subscription.
	subscription := models.Subscription{
		UserID:   user.ID,
		PlanID:   plan.ID,
		Status:   models.SubscriptionActive,
		StartsAt: time.Now(),
	}

	if err := db.Create(&subscription).Error; err != nil {
		t.Fatalf("create subscription: %v", err)
	}

	usageService := monitoring.NewUsageService(db)
	subscriptionService := NewService(db)

	checker := NewLimitChecker(
		subscriptionService,
		usageService,
	)

	// First request: should be allowed.
	if err := checker.CheckDailyMessageLimit(user.ID); err != nil {
		t.Fatalf("expected request to be allowed, got: %v", err)
	}

	// Record two successful messages.
	usages := []models.Usage{
		{
			UserID:    user.ID,
			Feature:   "conversation",
			Model:     "qwen3.5:4b",
			Status:    models.UsageSuccess,
			CreatedAt: time.Now().UTC(),
		},
		{
			UserID:    user.ID,
			Feature:   "reading",
			Model:     "qwen3.5:4b",
			Status:    models.UsageSuccess,
			CreatedAt: time.Now().UTC(),
		},
	}

	if err := db.Create(&usages).Error; err != nil {
		t.Fatalf("create usage records: %v", err)
	}

	// The daily limit is now reached.
	err := checker.CheckDailyMessageLimit(user.ID)

	if err == nil {
		t.Fatal("expected daily message limit error")
	}

	t.Logf("received expected limit error: %v", err)
}

func TestCheckMonthlyTokenLimit(t *testing.T) {
	db := setupTestDB(t)

	user := models.User{
		FullName:     "Limit Test User",
		Email:        "limit-test@example.com",
		PasswordHash: "test-hash",
		Role:         models.RoleUser,
		IsActive:     true,
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Create user: %v", err)
	}

	plan := models.Plan{
		Code:              models.PlanFree,
		Name:              "free",
		DailyMessageLimit: 2,
		MonthlyTokenLimit: 2,
		IsActive:          true,
	}

	if err := db.Create(&plan).Error; err != nil {
		t.Fatalf("create plan: %v", err)
	}

	usageService := monitoring.NewUsageService(db)
	subscriptionService := NewService(db)

	_, err := subscriptionService.CreateSubscription(
		user.ID,
		models.PlanFree,
	)
	if err != nil {
		t.Fatalf("create subscription: %v", err)
	}

	checker := NewLimitChecker(subscriptionService, usageService)

	if err := checker.CheckMonthlyTokenLimit(user.ID); err != nil {
		t.Fatalf("expected request to be allowed, got: %v", err)
	}

	usages := []models.Usage{
		{
			UserID:       user.ID,
			Feature:      "conversation",
			Model:        "qwen3.5:4b",
			InputTokens:  1,
			OutputTokens: 0,
			Status:       models.UsageSuccess,
			CreatedAt:    time.Now().UTC(),
		},
		{
			UserID:       user.ID,
			Feature:      "reading",
			Model:        "qwen3.5:4b",
			InputTokens:  1,
			OutputTokens: 0,
			Status:       models.UsageSuccess,
			CreatedAt:    time.Now().UTC(),
		},
	}

	if err := db.Create(&usages).Error; err != nil {
		t.Fatalf("create usage records: %v", err)
	}

	err = checker.CheckMonthlyTokenLimit(user.ID)

	if err == nil {
		t.Fatal("expected monthly token limit error")
	}

	t.Logf("received expected limit error: %v", err)
}
