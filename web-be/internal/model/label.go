package model

import (
	"time"

	"gorm.io/gorm"
)

// Label represents a label/tag for tasks.
type Label struct {
	ID        string         `gorm:"primaryKey;type:uuid"`
	Name      string         `gorm:"not null"`
	Color     string         `gorm:"size:7"`
	UserID    string         `gorm:"type:uuid;index;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the table name.
func (Label) TableName() string {
	return "labels"
}
