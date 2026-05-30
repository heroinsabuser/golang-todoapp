package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
	core_postgres_pool "github.com/heroinsabuser/golang-todoapp/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.todos(title, description, completed, created_at, completed_at, user_id) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, title, description, completed, created_at, completed_at, user_id;
	`

	row := r.pool.QueryRow(ctx, query, task.Title, task.Description, task.Completed, task.CreatedAt, task.CompletedAt, task.UserID)
	var taskModel TaskModel
	if err := row.Scan(&taskModel.ID, &taskModel.Version, &taskModel.Title, &taskModel.Description,
		&taskModel.Completed, &taskModel.CreatedAt, &taskModel.CompletedAt, &taskModel.UserID); err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf("user id key: %w", core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("scan row failed: %v: %w", task, err)
	}
	taskDomain := domain.NewTask(taskModel.ID, taskModel.Version, taskModel.Title, taskModel.Description, taskModel.Completed,
		taskModel.CreatedAt, taskModel.CompletedAt, taskModel.UserID)
	return taskDomain, nil
}
