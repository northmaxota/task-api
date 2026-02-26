package task

import "context"

type Storage interface {
	Create(ctx context.Context, task Task) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
}
