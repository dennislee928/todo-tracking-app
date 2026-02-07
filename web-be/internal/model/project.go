package model

import (
	"time"

	"gorm.io/gorm"
)

// Project represents a project/list.
type Project struct {
	ID        string         `gorm:"primaryKey;type:uuid"`
	Name      string         `gorm:"not null"`
	Color     string         `gorm:"size:7"`
	UserID    string         `gorm:"type:uuid;index;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Tasks   []Task          `gorm:"foreignKey:ProjectID"`
	Members []ProjectMember `gorm:"foreignKey:ProjectID"`
}

// TableName overrides the table name.
func (Project) TableName() string {
	return "projects"
}
