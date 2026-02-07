package dto

import (
	"time"

	"github.com/todo-tracking-app/web-be/internal/model"
)

// UserToVO converts model.User to UserVO.
func UserToVO(u model.User) UserVO {
	vo := UserVO{
		ID:        u.ID,
		Email:     u.Email,
		IsPremium: u.IsPremium,
	}
	if u.PremiumExpiresAt != nil {
		s := u.PremiumExpiresAt.Format(time.RFC3339)
		vo.PremiumExpiresAt = &s
	}
	return vo
}
