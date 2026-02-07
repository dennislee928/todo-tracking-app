package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"github.com/todo-tracking-app/web-be/api/grpc/proto"
	"github.com/todo-tracking-app/web-be/internal/model"
	"github.com/google/uuid"
)

// Server implements proto.TodoServiceServer.
type Server struct {
	proto.UnimplementedTodoServiceServer
	db *gorm.DB
}

// NewServer creates a new gRPC server.
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

// Serve starts the gRPC server on the given address.
func Serve(db *gorm.DB, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	proto.RegisterTodoServiceServer(s, NewServer(db))
	reflection.Register(s)
	log.Printf("gRPC server listening on %s", addr)
	return s.Serve(lis)
}

func (s *Server) ListTasks(ctx context.Context, req *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	q := s.db.Where("user_id = ?", req.UserId)
	if req.ProjectId != "" {
		q = q.Where("project_id = ?", req.ProjectId)
	}
	var tasks []model.Task
	if err := q.Find(&tasks).Error; err != nil {
		return nil, err
	}
	out := make([]*proto.TaskMessage, len(tasks))
	for i, t := range tasks {
		out[i] = taskToProto(&t)
	}
	return &proto.ListTasksResponse{Tasks: out}, nil
}

func (s *Server) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskMessage, error) {
	var task model.Task
	if err := s.db.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&task).Error; err != nil {
		return nil, err
	}
	return taskToProto(&task), nil
}

func (s *Server) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.TaskMessage, error) {
	task := model.Task{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Description: req.Description,
		ProjectID:  req.ProjectId,
		UserID:     req.UserId,
		Priority:   int(req.Priority),
		Status:     model.TaskStatusPending,
	}
	if req.DueDate != "" {
		if t, err := time.Parse(time.RFC3339, req.DueDate); err == nil {
			task.DueDate = &t
		}
	}
	if req.ReminderAt != "" {
		if t, err := time.Parse(time.RFC3339, req.ReminderAt); err == nil {
			task.ReminderAt = &t
		}
	}
	if err := s.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return taskToProto(&task), nil
}

func (s *Server) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.TaskMessage, error) {
	var task model.Task
	if err := s.db.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&task).Error; err != nil {
		return nil, err
	}
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.ProjectId != nil {
		task.ProjectID = *req.ProjectId
	}
	if req.Priority != nil {
		task.Priority = int(*req.Priority)
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if req.Progress != nil {
		task.Progress = int(*req.Progress)
	}
	if req.DueDate != nil && *req.DueDate != "" {
		if t, err := time.Parse(time.RFC3339, *req.DueDate); err == nil {
			task.DueDate = &t
		}
	}
	if req.ReminderAt != nil && *req.ReminderAt != "" {
		if t, err := time.Parse(time.RFC3339, *req.ReminderAt); err == nil {
			task.ReminderAt = &t
		}
	}
	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}
	return taskToProto(&task), nil
}

func (s *Server) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	s.db.Where("id = ? AND user_id = ?", req.Id, req.UserId).Delete(&model.Task{})
	return &proto.DeleteTaskResponse{}, nil
}

func (s *Server) ListProjects(ctx context.Context, req *proto.ListProjectsRequest) (*proto.ListProjectsResponse, error) {
	var projects []model.Project
	if err := s.db.Where("user_id = ?", req.UserId).Find(&projects).Error; err != nil {
		return nil, err
	}
	out := make([]*proto.ProjectMessage, len(projects))
	for i, p := range projects {
		out[i] = projectToProto(&p)
	}
	return &proto.ListProjectsResponse{Projects: out}, nil
}

func (s *Server) GetProject(ctx context.Context, req *proto.GetProjectRequest) (*proto.ProjectMessage, error) {
	var proj model.Project
	if err := s.db.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&proj).Error; err != nil {
		return nil, err
	}
	return projectToProto(&proj), nil
}

func (s *Server) CreateProject(ctx context.Context, req *proto.CreateProjectRequest) (*proto.ProjectMessage, error) {
	proj := model.Project{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Color:  req.Color,
		UserID: req.UserId,
	}
	if err := s.db.Create(&proj).Error; err != nil {
		return nil, err
	}
	return projectToProto(&proj), nil
}

func (s *Server) UpdateProject(ctx context.Context, req *proto.UpdateProjectRequest) (*proto.ProjectMessage, error) {
	var proj model.Project
	if err := s.db.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&proj).Error; err != nil {
		return nil, err
	}
	if req.Name != nil {
		proj.Name = *req.Name
	}
	if req.Color != nil {
		proj.Color = *req.Color
	}
	if err := s.db.Save(&proj).Error; err != nil {
		return nil, err
	}
	return projectToProto(&proj), nil
}

func (s *Server) DeleteProject(ctx context.Context, req *proto.DeleteProjectRequest) (*proto.DeleteProjectResponse, error) {
	s.db.Where("id = ? AND user_id = ?", req.Id, req.UserId).Delete(&model.Project{})
	return &proto.DeleteProjectResponse{}, nil
}

func taskToProto(t *model.Task) *proto.TaskMessage {
	m := &proto.TaskMessage{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		ProjectId:   t.ProjectID,
		UserId:      t.UserID,
		Priority:     int32(t.Priority),
		Status:      t.Status,
		Progress:    int32(t.Progress),
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   t.UpdatedAt.Format(time.RFC3339),
	}
	if t.DueDate != nil {
		m.DueDate = t.DueDate.Format(time.RFC3339)
	}
	if t.ReminderAt != nil {
		m.ReminderAt = t.ReminderAt.Format(time.RFC3339)
	}
	return m
}

func projectToProto(p *model.Project) *proto.ProjectMessage {
	return &proto.ProjectMessage{
		Id:        p.ID,
		Name:      p.Name,
		Color:     p.Color,
		UserId:    p.UserID,
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
	}
}
