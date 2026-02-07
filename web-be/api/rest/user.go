package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/todo-tracking-app/web-be/internal/dto"
	"github.com/todo-tracking-app/web-be/internal/model"
)

// RegisterUserRoutes registers user-related routes (require auth).
func RegisterUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &userHandler{db: db}
	r.GET("/me", h.GetMe)
}

type userHandler struct {
	db *gorm.DB
}

// GetMe returns the current authenticated user's profile.
// @Summary Get current user
// @Tags user
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.UserVO
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /me [get]
func (h *userHandler) GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var user model.User
	if err := h.db.Where("id = ?", userID.(string)).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, dto.UserToVO(user))
}
