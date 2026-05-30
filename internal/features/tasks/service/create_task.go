package tasks_service

import (
	"context"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
)

func (s *TasksService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("failed validate task: %w", err)
	}
	task, err := s.tasksRepository.CreateTask(ctx, task)

	if err != nil {
		return domain.Task{}, fmt.Errorf("create task failed: %w", err)
	}
	return task, nil
}
