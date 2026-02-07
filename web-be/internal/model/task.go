package model

import (
	"time"

	"gorm.io/gorm"
)

// Task status constants.
const (
	TaskStatusPending    = "pending"
	TaskStatusInProgress = "in_progress"
	TaskStatusCompleted   = "completed"
	TaskStatusCancelled   = "cancelled"
)

// Task represents a todo task.
type Task struct {
	ID          string          `gorm:"primaryKey;type:uuid"`
	Title       string          `gorm:"not null"`
	Description string          `gorm:"type:text"`
	ProjectID   string          `gorm:"type:uuid;index"`
	UserID      string          `gorm:"type:uuid;index;not null"`
	Priority    int             `gorm:"default:0"` // 0=none, 1=p4, 2=p3, 3=p2, 4=p1
	Status      string          `gorm:"size:20;default:pending"`
	DueDate     *time.Time      `gorm:"type:timestamptz"`
	ReminderAt  *time.Time      `gorm:"type:timestamptz"`
	Progress    int             `gorm:"default:0"` // 0-100
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`

	Project   *Project    `gorm:"foreignKey:ProjectID"`
	Labels    []Label     `gorm:"many2many:task_labels;"`
	Subtasks  []Subtask   `gorm:"foreignKey:TaskID"`
}

// TableName overrides the table name.
func (Task) TableName() string {
	return "tasks"
}
