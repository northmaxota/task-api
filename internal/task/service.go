package task

import "errors"

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Create(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("title is required")
	}

	task := Task{
		Title: title,
		Done:  false,
	}

	return s.storage.Create(task)
}

func (s *Service) GetAll() ([]Task, error) {
	return s.storage.GetAll()
}
