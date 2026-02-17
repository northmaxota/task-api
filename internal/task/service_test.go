package task_test

import (
	"testing"

	"github.com/northmaxota/task-api/internal/domain"
	"github.com/northmaxota/task-api/internal/task"
)

type fakeStorage struct {
	tasks []task.Task
}

func (f *fakeStorage) Create(t task.Task) (task.Task, error) {
	f.tasks = append(f.tasks, t)
	return t, nil
}

func (f *fakeStorage) GetAll() ([]task.Task, error) {
	return f.tasks, nil
}

func TestService_Create(t *testing.T) {
	storage := &fakeStorage{}
	svc := task.NewService(storage)

	taskTitle := "Test Task"
	createdTask, err := svc.Create(taskTitle)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createdTask.Title != taskTitle {
		t.Errorf("Expected title %q, got %q", taskTitle, createdTask.Title)
	}

	if createdTask.Done {
		t.Errorf("Expected Done to be false, got true")
	}
}

func TestService_Create_EmptyTitle(t *testing.T) {
	storage := &fakeStorage{}
	svc := task.NewService(storage)

	_, err := svc.Create("")
	if err == nil {
		t.Fatal("Expected error for empty title, got nil")
	}

	expectedErr := domain.ErrTitleRequired.Error()
	if err.Error() != expectedErr {
		t.Errorf("Expected error %q, got %q", expectedErr, err.Error())
	}
}

func TestService_GetAll(t *testing.T) {
	storage := &fakeStorage{}
	svc := task.NewService(storage)

	// Create some tasks
	svc.Create("Task 1")
	svc.Create("Task 2")

	tasks, err := svc.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	expectedTitles := []string{"Task 1", "Task 2"}
	for i, task := range tasks {
		if task.Title != expectedTitles[i] {
			t.Errorf("Expected title %q, got %q", expectedTitles[i], task.Title)
		}
	}
}
