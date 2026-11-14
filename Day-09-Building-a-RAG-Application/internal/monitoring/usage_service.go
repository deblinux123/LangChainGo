package monitoring

import (
	"fmt"

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

func (s *UsageService) RecordUsage(usage *models.Usage) error {
	if usage == nil {
		return fmt.Errorf("❌ usage cannot be nil")
	}

	if err := s.db.Create(usage).Error; err != nil {
		return fmt.Errorf("❌ record usage: %w", err)
	}

	return nil
}
