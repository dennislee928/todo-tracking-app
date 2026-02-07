package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	ID               string         `gorm:"primaryKey;type:uuid"`
	Email            string         `gorm:"uniqueIndex;not null"`
	Password         string         `gorm:"not null"`
	IsPremium        bool           `gorm:"not null;default:false"`
	PremiumExpiresAt *time.Time     `gorm:"type:timestamptz"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the table name.
func (User) TableName() string {
	return "users"
}
