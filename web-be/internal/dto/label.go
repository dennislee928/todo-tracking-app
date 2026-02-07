package dto

// LabelCreateRequest is the request body for creating a label.
type LabelCreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}

// LabelUpdateRequest is the request body for updating a label.
type LabelUpdateRequest struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

// LabelVO is the view object for label.
type LabelVO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	UserID string `json:"user_id"`
}
