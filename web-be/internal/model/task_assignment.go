package model

import (
	"time"
)

// TaskAssignment represents a task assignment to a user.
type TaskAssignment struct {
	TaskID    string    `gorm:"primaryKey;type:uuid"`
	UserID    string    `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// TableName overrides the table name.
func (TaskAssignment) TableName() string {
	return "task_assignments"
}
