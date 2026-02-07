package dto

// RegisterRequest is the request body for user registration.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest is the request body for user login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse is the response for auth endpoints.
type AuthResponse struct {
	Token string  `json:"token"`
	User  UserVO  `json:"user"`
}
