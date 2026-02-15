package task

import (
	"log"
	"time"
)

type loggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &loggingService{next: next}
}

func (l *loggingService) Create(title string) (Task, error) {
	task, err := l.next.Create(title)
	if err != nil {
		return Task{}, err
	}
	now := time.Now()
	log.Printf("Created task: %s at %s", task.Title, now.Format(time.RFC3339))
	return task, nil
}

func (l *loggingService) GetAll() ([]Task, error) {
	tasks, err := l.next.GetAll()
	if err != nil {
		return nil, err
	}
	log.Printf("Retrieved %d tasks at %s", len(tasks), time.Now().Format(time.RFC3339))
	return tasks, nil
}
