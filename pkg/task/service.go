package task

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, task Task) (string, error)
	GetTaskByID(ctx context.Context, taskID string) (*Task, error)
	DeleteTaskByID(ctx context.Context, taskID string) error
}

type taskService struct {
	tasks map[string]*Task
	mu sync.RWMutex
}

func NewService(tasks map[string]*Task) TaskService {
	return &taskService{tasks: tasks}
}

func (s *taskService) CreateTask(ctx context.Context, task Task) (string, error) {
	taskID := uuid.NewString()
	task.ID = taskID
	s.tasks[taskID] = &task

	go s.simulateWorker(ctx, &task)

	return taskID, nil
}

func (s *taskService) DeleteTaskByID(ctx context.Context, task_id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[task_id]
	if !exists {
		return fmt.Errorf("Task not found")
	}

	delete(s.tasks, task_id)
	return nil
}

func (s *taskService) GetTaskByID(ctx context.Context, task_id string) (*Task, error) {
	task, exists := s.tasks[task_id]
	if !exists {
		return &Task{}, fmt.Errorf("Task not found")
	}

	return task, nil
}

// simulateWorker имитирует длительную I/O bound операцию (например, 3-5 минут)
func (s *taskService) simulateWorker(ctx context.Context, task *Task) {
	s.mu.Lock()
	task.Status = StatusRunning
	task.StartedAt = time.Now().UTC()
	s.mu.Unlock()

	time.Sleep(3 * time.Minute)

	s.mu.Lock()
	task.FinishedAt = time.Now().UTC()
	task.Duration = task.FinishedAt.Sub(task.StartedAt)
	task.Status = StatusCompleted
	task.Result = "Mock данные для результата"
	s.mu.Unlock()
}

