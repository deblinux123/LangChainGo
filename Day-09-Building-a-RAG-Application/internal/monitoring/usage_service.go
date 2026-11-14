package monitoring

import (
	"fmt"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"gorm.io/gorm"
)

type UsageService struct {
	db *gorm.DB
}

func NewUsageService(db *gorm.DB) *UsageService {
	return &UsageService{
		db: db,
	}
}

// RecordUsage stores the usage information of an AI request.
func (s *UsageService) RecordUsage(usage *models.Usage) error {
	if usage == nil {
		return fmt.Errorf("usage cannot be nil")
	}

	if err := s.db.Create(usage).Error; err != nil {
		return fmt.Errorf("record usage: %w", err)
	}

	return nil
}

// GetDailyMessageCount returns the number of successful requests today.
func (s *UsageService) GetDailyMessageCount(userID uint) (int64, error) {
	now := time.Now().UTC()

	startOfDay := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, time.UTC,
	)

	startOfTomorrow := startOfDay.AddDate(0, 0, 1)

	var count int64

	err := s.db.Model(&models.Usage{}).
		Where(
			"user_id = ? AND status = ? AND created_at >= ? AND created_at < ?",
			userID,
			models.UsageSuccess,
			startOfDay,
			startOfTomorrow,
		).
		Count(&count).Error

	if err != nil {
		return 0, fmt.Errorf("get daily message count: %w", err)
	}

	return count, nil
}

// GetMonthlyTokenUsage returns the total successful tokens used this month.
func (s *UsageService) GetMonthlyTokenUsage(userID uint) (int64, error) {
	now := time.Now().UTC()

	startOfMonth := time.Date(
		now.Year(), now.Month(), 1,
		0, 0, 0, 0, time.UTC,
	)

	startOfNextMonth := startOfMonth.AddDate(0, 1, 0)

	var total int64

	err := s.db.Model(&models.Usage{}).
		Select("COALESCE(SUM(input_tokens + output_tokens), 0)").
		Where(
			"user_id = ? AND status = ? AND created_at >= ? AND created_at < ?",
			userID,
			models.UsageSuccess,
			startOfMonth,
			startOfNextMonth,
		).
		Scan(&total).Error

	if err != nil {
		return 0, fmt.Errorf("get monthly token usage: %w", err)
	}

	return total, nil
}
