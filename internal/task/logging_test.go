package task_test

import (
	"context"
	"testing"

	"github.com/northmaxota/task-api/internal/domain"
	"github.com/northmaxota/task-api/internal/task"
)

type fakeService struct {
	tasks     task.Task
	called    bool
	returnErr error
}

func (f *fakeService) Create(ctx context.Context, title string) (task.Task, error) {
	f.called = true
	f.tasks = task.Task{Title: title}
	return f.tasks, f.returnErr
}

func (f *fakeService) GetAll(ctx context.Context) ([]task.Task, error) {
	return nil, nil
}

func TestLoggingService_Create_Success(t *testing.T) {
	fake := &fakeService{}
	loggingSvc := task.NewLoggingService(fake)
	_, err := loggingSvc.Create(t.Context(), "Test Task")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !fake.called {
		t.Error("expected Create to be called on the next service")
	}
}

func TestLoggingService_Create_Error(t *testing.T) {
	fake := &fakeService{returnErr: domain.ErrTitleRequired}
	loggingSvc := task.NewLoggingService(fake)
	_, err := loggingSvc.Create(t.Context(), "")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
	if !fake.called {
		t.Error("expected Create to be called on the next service")
	}
}
