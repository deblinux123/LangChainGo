package subscription

import (
	"fmt"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/monitoring"
)

type LimitChecker struct {
	subscriptionService *Service
	usageSevice         *monitoring.UsageService
}

func NewLimitChecker(
	subscriptionService *Service,
	usageService *monitoring.UsageService,
) *LimitChecker {
	return &LimitChecker{
		subscriptionService: subscriptionService,
		usageSevice:         usageService,
	}
}

func (c *LimitChecker) CheckDailyMessageLimit(userID uint) error {
	subscription, err := c.subscriptionService.GetActiveSubscription(userID)

	if err != nil {
		return fmt.Errorf("❌ check subscription: %w", err)
	}

	messageCount, err := c.usageSevice.GetDailyMessageCount(userID)

	if err != nil {
		return fmt.Errorf("❌ check daily usage: %w", err)
	}

	limit := int64(subscription.Plan.DailyMessageLimit)

	if messageCount >= limit {
		return fmt.Errorf(
			"🔴 daily message limit reached: used %d of %d messages",
			messageCount,
			limit,
		)
	}

	return nil
}

func (c *LimitChecker) CheckMonthlyTokenLimit(userID uint) error {
	subscription, err := c.subscriptionService.GetActiveSubscription(userID)

	if err != nil {
		return fmt.Errorf("❌ check subscription: %w", err)
	}

	messageCount, err := c.usageSevice.GetMonthlyTokenUsage(userID)

	if err != nil {
		return fmt.Errorf("❌ check monthly usage: %w", err)
	}

	limit := int64(subscription.Plan.MonthlyTokenLimit)

	if messageCount >= limit {
		return fmt.Errorf(
			"🔴 Monthly token limit reached: used %d of %d messages",
			messageCount,
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
