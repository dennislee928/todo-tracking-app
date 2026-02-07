package model

import (
	"time"
)

// ProjectMember represents a project sharing/collaboration.
type ProjectMember struct {
	ProjectID string    `gorm:"primaryKey;type:uuid"`
	UserID    string    `gorm:"primaryKey;type:uuid"`
	Role      string    `gorm:"size:20;default:member"` // owner, admin, member
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// TableName overrides the table name.
func (ProjectMember) TableName() string {
	return "project_members"
}
