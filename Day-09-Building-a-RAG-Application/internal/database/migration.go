package database

import (
	"fmt"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Plan{},
		&models.Subscription{},
		&models.Usage{},
	); err != nil {
		return fmt.Errorf("❌ database migration failed: %w", err)
	}

	return nil
}
