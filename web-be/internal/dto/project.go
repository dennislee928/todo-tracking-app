package dto

import "time"

// ProjectCreateRequest is the request body for creating a project.
type ProjectCreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}

// ProjectUpdateRequest is the request body for updating a project.
type ProjectUpdateRequest struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

// ProjectVO is the view object for project.
type ProjectVO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
