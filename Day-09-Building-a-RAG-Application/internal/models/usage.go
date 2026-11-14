package models

import "time"

type UsageStatus string

const (
	UsageSuccess UsageStatus = "success"
	UsageFailed  UsageStatus = "failed"
)

type Usage struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `gorm:"not null;index" json:"user_id"`

	Feature string `gorm:"size:50;not null;index" json:"feature"`

	Model string `gorm:"size:100;not null" json:"model"`

	InputTokens  int64 `gorm:"not null;default:0" json:"input_tokens"`
	OutputTokens int64 `gorm:"not null;default:0" json:"output_tokens"`

	LatencyMs int64 `gorm:"not null;default:0" json:"latency_ms"`

	Status UsageStatus `gorm:"size:20;not null;index" json:"status"`

	CreatedAt time.Time `gorm:"index" json:"created_at"`
}
