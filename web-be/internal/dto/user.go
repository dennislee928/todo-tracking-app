package dto

// UserVO is the view object for user.
type UserVO struct {
	ID               string  `json:"id"`
	Email            string  `json:"email"`
	IsPremium        bool    `json:"is_premium"`
	PremiumExpiresAt *string `json:"premium_expires_at,omitempty"`
}
