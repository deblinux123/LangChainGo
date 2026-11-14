package models

import "time"

type PlanCode string

const (
	PlanFree    PlanCode = "free"
	PlanPremium PlanCode = "premium"
)

type Plan struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	Code                PlanCode       `gorm:"size:30;uniqueIndex;not null" json:"code"`
	Name                string         `gorm:"size:100;not null" json:"name"`
	DailyMessageLimit   int            `gorm:"not null;default:20" json:"daily_message_limit"`
	MonthlyTokenLimit   int64          `gorm:"not null;default:100000" json:"monthly_token_limit"`
	CanAccessAllLessons bool           `gorm:"not null;default:false" json:"can_access_all_lessons"`
	CanAccessAdvanced   bool           `gorm:"not null;default:false" json:"can_access_advanced"`
	IsActive            bool           `gorm:"not null;default:true" json:"is_active"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	Subscriptions       []Subscription `gorm:"foreignKey:PlanID" json:"-"`
}
