package subscription

import (
	"fmt"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/monitoring"
)

type LimitChecker struct {
	subscriptionService *Service
	usageService        *monitoring.UsageService
}

func NewLimitChecker(
	subscriptionService *Service,
	usageService *monitoring.UsageService,
) *LimitChecker {
	return &LimitChecker{
		subscriptionService: subscriptionService,
		usageService:        usageService,
	}
}

func (c *LimitChecker) CheckDailyMessageLimit(userID uint) error {
	subscription, err := c.subscriptionService.GetActiveSubscription(userID)
	if err != nil {
		return fmt.Errorf("check subscription: %w", err)
	}

	messageCount, err := c.usageService.GetDailyMessageCount(userID)
	if err != nil {
		return fmt.Errorf("check daily usage: %w", err)
	}

	limit := int64(subscription.Plan.DailyMessageLimit)

	if messageCount >= limit {
		return fmt.Errorf(
			"daily message limit reached: used %d of %d messages",
			messageCount,
			limit,
		)
	}

	return nil
}

func (c *LimitChecker) CheckMonthlyTokenLimit(userID uint) error {
	subscription, err := c.subscriptionService.GetActiveSubscription(userID)
	if err != nil {
		return fmt.Errorf("check subscription: %w", err)
	}

	tokenUsage, err := c.usageService.GetMonthlyTokenUsage(userID)
	if err != nil {
		return fmt.Errorf("check monthly usage: %w", err)
	}

	limit := subscription.Plan.MonthlyTokenLimit

	if tokenUsage >= limit {
		return fmt.Errorf(
			"monthly token limit reached: used %d of %d tokens",
			tokenUsage,
			limit,
		)
	}

	return nil
}

func (c *LimitChecker) CheckLimits(userID uint) error {
	if err := c.CheckDailyMessageLimit(userID); err != nil {
		return err
	}

	if err := c.CheckMonthlyTokenLimit(userID); err != nil {
		return err
	}

	return nil
}
