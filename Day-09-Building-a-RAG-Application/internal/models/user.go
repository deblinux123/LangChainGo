package models

import "time"

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	FullName      string         `grom:"size:150;not null" json:"full_name"`
	Email         string         `gorm:"size:255;uniqeIndex;not null" json:"email"`
	PasswordHash  string         `gorm:"not null" json:"-"`
	Role          UserRole       `gorm:"size:20;not null;default:user" json:"role"`
	IsActive      bool           `gorm:"not null;default:true" json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     time.Time      `gorm:"index" json:"-"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID" json:"-"`
}
