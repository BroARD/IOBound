package task

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Мок-сервис для теста
type mockTaskService struct{}

func (m *mockTaskService) CreateTask(ctx context.Context, task Task) (string, error) {
    return "mock-task-id", nil
}

func (m *mockTaskService) DeleteTaskByID(ctx context.Context, task_id string) error {
    return nil
}

func (m *mockTaskService) GetTaskByID(ctx context.Context, task_id string) (*Task, error) {
    return &Task{}, nil
}


func TestCreateTaskHandler(t *testing.T) {
    h := &taskHandler{
        service: &mockTaskService{},
    }

    req := httptest.NewRequest("POST", "/tasks", nil)
    w := httptest.NewRecorder()

    h.CreateTask(w, req)

    resp := w.Result()
    assert.Equal(t, http.StatusCreated, resp.StatusCode)

    body := w.Body.String()
    assert.Equal(t, "mock-task-id", body)
}

func TestDeleteTaskByID(t *testing.T) {
    h := &taskHandler{
        service: &mockTaskService{},
    }

    req := httptest.NewRequest("DELETE", "/tasks", nil)
    w := httptest.NewRecorder()

    h.DeleteTaskByID(w, req)

    resp := w.Result()
    assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetTaskByID(t *testing.T) {
    h := &taskHandler{
        service: &mockTaskService{},
    }

    req := httptest.NewRequest("GET", "/tasks", nil)
    w := httptest.NewRecorder()

    h.GetTaskByID(w, req)

    resp := w.Result()
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    
}
