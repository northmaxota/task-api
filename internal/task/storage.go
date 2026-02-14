package task

type Storage interface {
	Create(task Task) (Task, error)
	GetAll() ([]Task, error)
}
