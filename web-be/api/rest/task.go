package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/todo-tracking-app/web-be/internal/dto"
	"github.com/todo-tracking-app/web-be/internal/model"
	"github.com/jinzhu/copier"
)

// RegisterTaskRoutes registers task routes.
func RegisterTaskRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &taskHandler{db: db}
	tasks := r.Group("/tasks")
	{
		tasks.GET("", h.List)
		tasks.GET("/today", h.Today)
		tasks.GET("/upcoming", h.Upcoming)
		tasks.POST("", h.Create)
		tasks.GET("/:id", h.Get)
		tasks.PUT("/:id", h.Update)
		tasks.DELETE("/:id", h.Delete)
	}
}

type taskHandler struct {
	db *gorm.DB
}

func (h *taskHandler) getUserID(c *gin.Context) string {
	uid, _ := c.Get("user_id")
	return uid.(string)
}

// List returns tasks (optionally filtered by project).
// @Summary List tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param project_id query string false "Filter by project ID"
// @Success 200 {array} dto.TaskVO
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (h *taskHandler) List(c *gin.Context) {
	userID := h.getUserID(c)
	projectID := c.Query("project_id")

	q := h.db.Where("user_id = ?", userID)
	if projectID != "" {
		q = q.Where("project_id = ?", projectID)
	}
	var tasks []model.Task
	if err := q.Preload("Labels").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vos []dto.TaskVO
	for _, t := range tasks {
		vo := taskToVO(t)
		vos = append(vos, vo)
	}
	c.JSON(http.StatusOK, vos)
}

// Today returns tasks due today.
// @Summary List today's tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.TaskVO
// @Failure 500 {object} map[string]string
// @Router /tasks/today [get]
func (h *taskHandler) Today(c *gin.Context) {
	userID := h.getUserID(c)
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.Add(24 * time.Hour)

	var tasks []model.Task
	if err := h.db.Where("user_id = ? AND due_date >= ? AND due_date < ? AND status != ?",
		userID, start, end, model.TaskStatusCompleted).Preload("Labels").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vos []dto.TaskVO
	for _, t := range tasks {
		vos = append(vos, taskToVO(t))
	}
	c.JSON(http.StatusOK, vos)
}

// Upcoming returns tasks due in the future.
// @Summary List upcoming tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param days query int false "Number of days ahead" default(7)
// @Success 200 {array} dto.TaskVO
// @Failure 500 {object} map[string]string
// @Router /tasks/upcoming [get]
func (h *taskHandler) Upcoming(c *gin.Context) {
	userID := h.getUserID(c)
	days := 7
	if d := c.Query("days"); d != "" {
		if n, err := parseInt(d); err == nil && n > 0 {
			days = n
		}
	}
	now := time.Now()
	end := now.AddDate(0, 0, days)

	var tasks []model.Task
	if err := h.db.Where("user_id = ? AND due_date > ? AND due_date <= ? AND status != ?",
		userID, now, end, model.TaskStatusCompleted).Preload("Labels").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vos []dto.TaskVO
	for _, t := range tasks {
		vos = append(vos, taskToVO(t))
	}
	c.JSON(http.StatusOK, vos)
}

// Create creates a new task.
// @Summary Create task
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.TaskCreateRequest true "Task create request"
// @Success 201 {object} dto.TaskVO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (h *taskHandler) Create(c *gin.Context) {
	var req dto.TaskCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := model.Task{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		ProjectID:   req.ProjectID,
		UserID:      h.getUserID(c),
		Priority:    req.Priority,
		Status:      model.TaskStatusPending,
	}
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}
	if req.ReminderAt != nil {
		task.ReminderAt = req.ReminderAt
	}
	if err := h.db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.db.Preload("Labels").First(&task, "id = ?", task.ID)
	c.JSON(http.StatusCreated, taskToVO(task))
}

// Get returns a task by ID.
// @Summary Get task
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} dto.TaskVO
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func (h *taskHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	if err := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).Preload("Labels").Preload("Subtasks").First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, taskToVO(task))
}

// Update updates a task.
// @Summary Update task
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param body body dto.TaskUpdateRequest true "Task update request"
// @Success 200 {object} dto.TaskVO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [put]
func (h *taskHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.TaskUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task model.Task
	if err := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	copier.CopyWithOption(&task, &req, copier.Option{IgnoreEmpty: true})
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}
	if req.ReminderAt != nil {
		task.ReminderAt = req.ReminderAt
	}
	if err := h.db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.db.Preload("Labels").Preload("Subtasks").First(&task, "id = ?", task.ID)
	c.JSON(http.StatusOK, taskToVO(task))
}

// Delete deletes a task.
// @Summary Delete task
// @Tags tasks
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [delete]
func (h *taskHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	result := h.db.Where("id = ? AND user_id = ?", id, h.getUserID(c)).Delete(&model.Task{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func taskToVO(t model.Task) dto.TaskVO {
	vo := dto.TaskVO{}
	_ = copier.Copy(&vo, &t)
	for _, l := range t.Labels {
		vo.Labels = append(vo.Labels, dto.LabelVO{ID: l.ID, Name: l.Name, Color: l.Color})
	}
	return vo
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}
