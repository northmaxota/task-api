package task_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/northmaxota/task-api/internal/domain"
	"github.com/northmaxota/task-api/internal/task"
)

type fakeStorage struct {
	tasks []task.Task
}

func (f *fakeStorage) Create(ctx context.Context, t task.Task) (task.Task, error) {
	select {
	case <-time.After(100 * time.Millisecond):
		f.tasks = append(f.tasks, t)
		return t, nil
	case <-ctx.Done():
		return task.Task{}, ctx.Err()
	}
}

func (f *fakeStorage) GetAll(ctx context.Context) ([]task.Task, error) {
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
			_, err := svc.Create(context.Background(), tt.title)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %q, got %q", tt.wantErr, err)
			}
		})
	}
}

func TestService_Create_ContextCancelled(t *testing.T) {
	storage := &fakeStorage{}
	svc := task.NewService(storage)

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	_, err := svc.Create(ctx, "Test Task")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Expected error %q, got %q", context.Canceled, err)
	}
}

func TestService_GetAll(t *testing.T) {
	storage := &fakeStorage{}
	svc := task.NewService(storage)

	// Create some tasks
	_, err := svc.Create(t.Context(), "Task 1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = svc.Create(t.Context(), "Task 2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := svc.GetAll(t.Context())
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
