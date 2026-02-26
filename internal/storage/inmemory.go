package storage

import (
	"context"
	"sync"

	"github.com/northmaxota/task-api/internal/task"
)

type InMemoryStorage struct {
	mu     sync.Mutex
	tasks  []task.Task
	nextID int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		tasks:  make([]task.Task, 0),
		nextID: 1,
	}
}

func (s *InMemoryStorage) Create(ctx context.Context, task task.Task) (task.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s *InMemoryStorage) GetAll(ctx context.Context) ([]task.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]task.Task(nil), s.tasks...), nil
}
