package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/northmaxota/task-api/internal/domain"
	httpServer "github.com/northmaxota/task-api/internal/http"
	"github.com/northmaxota/task-api/internal/task"
)

type fakeService struct {
	createCalled bool
	getAllCalled bool
	returnTask   task.Task
	returnTasks  []task.Task
	returnErr    error
}

func (f *fakeService) Create(title string) (task.Task, error) {
	f.createCalled = true
	return f.returnTask, f.returnErr
}

func (f *fakeService) GetAll() ([]task.Task, error) {
	f.getAllCalled = true
	return f.returnTasks, f.returnErr
}

func TestRouter_CreateTask_Success(t *testing.T) {
	fake := &fakeService{
		returnTask: task.Task{
			ID:    1,
			Title: "Test Task",
			Done:  false,
		},
	}

	handler := httpServer.NewHandler(fake)
	router := httpServer.NewRouter(handler)

	body := `{"title":"Test Task"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	if !fake.createCalled {
		t.Fatal("expected service.Create to be called")
	}

	var resp task.Task
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Title != "Test Task" {
		t.Fatalf("unexpected title: %s", resp.Title)
	}
}

func TestRouter_CreateTask_ValidationError(t *testing.T) {
	fake := &fakeService{
		returnErr: domain.ErrTitleRequired,
	}

	handler := httpServer.NewHandler(fake)
	router := httpServer.NewRouter(handler)

	body := `{"title":""}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestRouter_GetTasks_Success(t *testing.T) {
	fake := &fakeService{
		returnTasks: []task.Task{
			{ID: 1, Title: "Task 1"},
			{ID: 2, Title: "Task 2"},
		},
	}

	handler := httpServer.NewHandler(fake)
	router := httpServer.NewRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	if !fake.getAllCalled {
		t.Fatal("expected service.GetAll to be called")
	}

	var resp []task.Task
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(resp))
	}
}
