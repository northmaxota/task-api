package task

import (
	"context"

	"github.com/northmaxota/task-api/internal/domain"
)

type TaskService struct {
	storage Storage
}

type Service interface {
	Create(ctx context.Context, title string) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
}

func NewService(storage Storage) Service {
	return &TaskService{storage: storage}
}

func (s *TaskService) Create(ctx context.Context, title string) (Task, error) {
	if title == "" {
		return Task{}, domain.ErrTitleRequired
	}

	task := Task{
		Title: title,
		Done:  false,
	}

	select {
	case <-ctx.Done():
		return Task{}, ctx.Err()
	default:
	}

	return s.storage.Create(ctx, task)
}

func (s *TaskService) GetAll(ctx context.Context) ([]Task, error) {
	return s.storage.GetAll(ctx)
}
