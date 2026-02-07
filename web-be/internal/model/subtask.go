package model

import (
	"time"

	"gorm.io/gorm"
)

// Subtask represents a sub-task under a task.
type Subtask struct {
	ID        string         `gorm:"primaryKey;type:uuid"`
	TaskID    string         `gorm:"type:uuid;index;not null"`
	Title     string         `gorm:"not null"`
	Completed bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the table name.
func (Subtask) TableName() string {
	return "subtasks"
}
