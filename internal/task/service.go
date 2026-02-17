package task

import (
	"github.com/northmaxota/task-api/internal/domain"
)

type TaskService struct {
	storage Storage
}

type Service interface {
	Create(title string) (Task, error)
	GetAll() ([]Task, error)
}

func NewService(storage Storage) Service {
	return &TaskService{storage: storage}
}

func (s *TaskService) Create(title string) (Task, error) {
	if title == "" {
		return Task{}, domain.ErrTitleRequired
	}

	task := Task{
		Title: title,
		Done:  false,
	}

	return s.storage.Create(task)
}

func (s *TaskService) GetAll() ([]Task, error) {
	return s.storage.GetAll()
}
