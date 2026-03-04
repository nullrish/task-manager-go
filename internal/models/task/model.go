package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID              uuid.UUID `json:"id" db:"id"`
	TaskTitle       string    `json:"task_title" db:"task_title"`
	TaskDescription string    `json:"task_description" db:"title_description"`
	UserID          uuid.UUID `json:"user_id" db:"user_id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type TaskRequest struct {
	ID              uuid.UUID `json:"id"`
	TaskTitle       string    `json:"task_title"`
	TaskDescription string    `json:"task_description"`
	UserID          uuid.UUID `json:"user_id"`
}
