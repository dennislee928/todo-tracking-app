package dto

import "time"

// TaskCreateRequest is the request body for creating a task.
type TaskCreateRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	ProjectID   string    `json:"project_id"`
	Priority    int       `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	ReminderAt  *time.Time `json:"reminder_at"`
	LabelIDs    []string  `json:"label_ids"`
}

// TaskUpdateRequest is the request body for updating a task.
type TaskUpdateRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	ProjectID   *string    `json:"project_id"`
	Priority    *int       `json:"priority"`
	Status      *string    `json:"status"`
	Progress    *int       `json:"progress"`
	DueDate     *time.Time `json:"due_date"`
	ReminderAt  *time.Time `json:"reminder_at"`
	LabelIDs    []string   `json:"label_ids"`
}

// TaskVO is the view object for task.
type TaskVO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ProjectID   string    `json:"project_id"`
	UserID      string    `json:"user_id"`
	Priority    int       `json:"priority"`
	Status      string    `json:"status"`
	Progress    int       `json:"progress"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	ReminderAt  *time.Time `json:"reminder_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Labels      []LabelVO `json:"labels,omitempty"`
}
