package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/todo-tracking-app/web-be/internal/dto"
	"github.com/todo-tracking-app/web-be/internal/model"
	"github.com/jinzhu/copier"
)

// RegisterLabelRoutes registers label routes.
func RegisterLabelRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &labelHandler{db: db}
	labels := r.Group("/labels")
	{
		labels.GET("", h.List)
		labels.POST("", h.Create)
		labels.GET("/:id", h.Get)
		labels.PUT("/:id", h.Update)
		labels.DELETE("/:id", h.Delete)
	}
}

type labelHandler struct {
	db *gorm.DB
}

func (h *labelHandler) getUserID(c *gin.Context) string {
	uid, _ := c.Get("user_id")
	return uid.(string)
}

// List returns all labels for the user.
// @Summary List labels
// @Tags labels
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.LabelVO
// @Failure 500 {object} map[string]string
// @Router /labels [get]
func (h *labelHandler) List(c *gin.Context) {
	userID := h.getUserID(c)
	var labels []model.Label
	if err := h.db.Where("user_id = ?", userID).Find(&labels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vos []dto.LabelVO
	_ = copier.Copy(&vos, &labels)
	c.JSON(http.StatusOK, vos)
}

// Create creates a new label.
// @Summary Create label
// @Tags labels
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.LabelCreateRequest true "Label create request"
// @Success 201 {object} dto.LabelVO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /labels [post]
func (h *labelHandler) Create(c *gin.Context) {
	var req dto.LabelCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	label := model.Label{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Color:  req.Color,
		UserID: h.getUserID(c),
	}
	if err := h.db.Create(&label).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vo dto.LabelVO
	_ = copier.Copy(&vo, &label)
	c.JSON(http.StatusCreated, vo)
}

// Get returns a label by ID.
// @Summary Get label
// @Tags labels
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Label ID"
// @Success 200 {object} dto.LabelVO
// @Failure 404 {object} map[string]string
// @Router /labels/{id} [get]
func (h *labelHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var label model.Label
	if err := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).First(&label).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "label not found"})
		return
	}
	var vo dto.LabelVO
	_ = copier.Copy(&vo, &label)
	c.JSON(http.StatusOK, vo)
}

// Update updates a label.
// @Summary Update label
// @Tags labels
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Label ID"
// @Param body body dto.LabelUpdateRequest true "Label update request"
// @Success 200 {object} dto.LabelVO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /labels/{id} [put]
func (h *labelHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.LabelUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var label model.Label
	if err := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).First(&label).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "label not found"})
		return
	}
	if req.Name != nil {
		label.Name = *req.Name
	}
	if req.Color != nil {
		label.Color = *req.Color
	}
	if err := h.db.Save(&label).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vo dto.LabelVO
	_ = copier.Copy(&vo, &label)
	c.JSON(http.StatusOK, vo)
}

// Delete deletes a label.
// @Summary Delete label
// @Tags labels
// @Security BearerAuth
// @Param id path string true "Label ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /labels/{id} [delete]
func (h *labelHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	result := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).Delete(&model.Label{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "label not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
