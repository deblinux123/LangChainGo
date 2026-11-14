package subscription

import (
	"fmt"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetPlanByCode(code models.PlanCode) (*models.Plan, error) {
	var plan models.Plan

	err := s.db.Where(
		"code = ? AND is_active = ?", code, true,
	).First(&plan).Error

	if err != nil {
		return nil, fmt.Errorf("❌ get plan %s: %w", code, err)
	}

	return &plan, nil
}

func (s *Service) CreateSubscription(
	userID uint,
	planCode models.PlanCode,
) (*models.Subscription, error) {
	var subscription models.Subscription

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var user models.User

		if err := tx.First(&user, userID).Error; err != nil {
			return fmt.Errorf("❌ find user %d: %w", userID, err)
		}

		var plan models.Plan

		if err := tx.Where(
			"code = ? AND is_active = ?", planCode, true,
		).First(&plan).Error; err != nil {
			return fmt.Errorf("❌ find plan %s: %w", planCode, err)
		}

		var count int64
		if err := tx.Model(&models.Subscription{}).
			Where("user_id = ? AND status = ?", userID, models.SubscriptionActive).
			Count(&count).Error; err != nil {
			return fmt.Errorf("check existing subscription: %w", err)
		}

		if count > 0 {
			return fmt.Errorf("user %d already has an active subscription", userID)
		}

		subscription = models.Subscription{
			UserID:   user.ID,
			PlanID:   plan.ID,
			Status:   models.SubscriptionActive,
			StartsAt: time.Now(),
			EndsAt:   nil,
		}

		if err := tx.Create(&subscription).Error; err != nil {
			return fmt.Errorf("❌ create subscription: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}
func (s *Service) GetActiveSubscription(userID uint) (*models.Subscription, error) {
	var subscription models.Subscription

	err := s.db.
		Where(
			"user_id = ? AND status = ? AND (ends_at IS NULL OR ends_at > ?)",
			userID,
			models.SubscriptionActive,
			time.Now(),
		).
		First(&subscription).Error

	if err != nil {
		return nil, fmt.Errorf(
			"get active subscription for user %d: %w",
			userID,
			err,
		)
	}

	var plan models.Plan

	err = s.db.
		Where("id = ?", subscription.PlanID).
		First(&plan).Error

	if err != nil {
		return nil, fmt.Errorf(
			"get plan %d for subscription %d: %w",
			subscription.PlanID,
			subscription.ID,
			err,
		)
	}

	subscription.Plan = plan

	return &subscription, nil
}
