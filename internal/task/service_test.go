package task_test

import (
	"errors"
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
	tests := []struct {
		name    string
		title   string
		wantErr error
	}{
		{
			name:    "Valid Title",
			title:   "Test Task",
			wantErr: nil,
		}, {
			name:    "Empty Title",
			title:   "",
			wantErr: domain.ErrTitleRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &fakeStorage{}
			svc := task.NewService(storage)
			_, err := svc.Create(tt.title)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %q, got %q", tt.wantErr, err)
			}
		})
	}
}

func TestService_GetAll(t *testing.T) {
	storage := &fakeStorage{}
	svc := task.NewService(storage)

	// Create some tasks
	_, err := svc.Create("Task 1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = svc.Create("Task 2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

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
