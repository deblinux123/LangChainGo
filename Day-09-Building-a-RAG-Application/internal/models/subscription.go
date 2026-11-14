package models

import "time"

type SubscriptionStatus string

const (
	SubscriptionActive   SubscriptionStatus = "active"
	SubscriptionCanceled SubscriptionStatus = "canceled"
	SubscriptionExpired  SubscriptionStatus = "expired"
	SubscriptionPending  SubscriptionStatus = "pending"
)

type Subscription struct {
	ID uint `gorm:"primarykey" json:"id"`

	UserID uint `gorm:"not null;index" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`

	PlanID uint `gorm:"not null;index" json:"plan_id"`
	Plan   Plan `gorm:"foreignKey:PlanID" json:"plan,omitempty"`

	Status SubscriptionStatus `gorm:"size:20;not null;index" json:"status"`

	StartAt time.Time  `gorm:"not null" json:"start_at"`
	EndAt   *time.Time `json:"end_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
