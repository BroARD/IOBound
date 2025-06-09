package task

import "time"

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusFailed    TaskStatus = "failed"
)

type Task struct {
	ID         string
	Status     TaskStatus
	Result     string
	CreatedAt  time.Time
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   time.Duration
}
