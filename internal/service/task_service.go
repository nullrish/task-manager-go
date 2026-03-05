package service

import (
	"github.com/nullrish/task-manager-go/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}
