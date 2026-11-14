package database

import (
	"fmt"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"gorm.io/gorm"
)

func SeedPlan(db *gorm.DB) error {
	plans := []models.Plan{
		{
			Code:                models.PlanFree,
			Name:                "Free",
			DailyMessageLimit:   20,
			MonthlyTokenLimit:   100000,
			CanAccessAllLessons: false,
			CanAccessAdvanced:   false,
			IsActive:            true,
		},
		{
			Code:                models.PlanPremium,
			Name:                "Premium",
			DailyMessageLimit:   1000,
			MonthlyTokenLimit:   3000000,
			CanAccessAllLessons: true,
			CanAccessAdvanced:   true,
			IsActive:            true,
		},
	}

	for _, plan := range plans {
		if err := db.Where(
			models.Plan{Code: plan.Code},
		).FirstOrCreate(&plan).Error; err != nil {
			return fmt.Errorf("seed plan %s: %w", plan.Code, err)
		}
	}

	return nil
}
