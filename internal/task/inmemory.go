package task

import "sync"

type InMemoryStorage struct {
	mu     sync.Mutex
	tasks  []Task
	nextID int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (s *InMemoryStorage) Create(task Task) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s *InMemoryStorage) GetAll() ([]Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]Task(nil), s.tasks...), nil
}
