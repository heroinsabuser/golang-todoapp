package tasks_service

import (
	"context"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

func NewTasksService(usersRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: usersRepository,
	}
}
