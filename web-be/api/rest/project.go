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

// RegisterProjectRoutes registers project routes.
func RegisterProjectRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &projectHandler{db: db}
	projects := r.Group("/projects")
	{
		projects.GET("", h.List)
		projects.POST("", h.Create)
		projects.GET("/:id", h.Get)
		projects.PUT("/:id", h.Update)
		projects.DELETE("/:id", h.Delete)
	}
}

type projectHandler struct {
	db *gorm.DB
}

func (h *projectHandler) getUserID(c *gin.Context) string {
	uid, _ := c.Get("user_id")
	return uid.(string)
}

// List returns all projects for the user.
// @Summary List projects
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.ProjectVO
// @Failure 500 {object} map[string]string
// @Router /projects [get]
func (h *projectHandler) List(c *gin.Context) {
	userID := h.getUserID(c)
	var projects []model.Project
	if err := h.db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vos []dto.ProjectVO
	_ = copier.Copy(&vos, &projects)
	c.JSON(http.StatusOK, vos)
}

// Create creates a new project.
// @Summary Create project
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.ProjectCreateRequest true "Project create request"
// @Success 201 {object} dto.ProjectVO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /projects [post]
func (h *projectHandler) Create(c *gin.Context) {
	var req dto.ProjectCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	proj := model.Project{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Color:  req.Color,
		UserID: h.getUserID(c),
	}
	if err := h.db.Create(&proj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vo dto.ProjectVO
	_ = copier.Copy(&vo, &proj)
	c.JSON(http.StatusCreated, vo)
}

// Get returns a project by ID.
// @Summary Get project
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Project ID"
// @Success 200 {object} dto.ProjectVO
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [get]
func (h *projectHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var proj model.Project
	if err := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).First(&proj).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	var vo dto.ProjectVO
	_ = copier.Copy(&vo, &proj)
	c.JSON(http.StatusOK, vo)
}

// Update updates a project.
// @Summary Update project
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Project ID"
// @Param body body dto.ProjectUpdateRequest true "Project update request"
// @Success 200 {object} dto.ProjectVO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [put]
func (h *projectHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.ProjectUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var proj model.Project
	if err := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).First(&proj).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	if req.Name != nil {
		proj.Name = *req.Name
	}
	if req.Color != nil {
		proj.Color = *req.Color
	}
	if err := h.db.Save(&proj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vo dto.ProjectVO
	_ = copier.Copy(&vo, &proj)
	c.JSON(http.StatusOK, vo)
}

// Delete deletes a project.
// @Summary Delete project
// @Tags projects
// @Security BearerAuth
// @Param id path string true "Project ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [delete]
func (h *projectHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	result := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).Delete(&model.Project{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
