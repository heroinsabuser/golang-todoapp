package domain

import (
	"fmt"
	"time"

	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	UserID      int
}

func (task *Task) Validate() error {
	titleLength := len([]rune(task.Title))
	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf("task title must be between 1 and 100: %w", core_errors.ErrInvalidArgument)
	}
	if task.Description != nil {
		descriptionLength := len([]rune(*task.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf("task description must be between 1 and 1000: %w", core_errors.ErrInvalidArgument)
		}
	}
	if task.Completed {
		if task.CompletedAt == nil {
			return fmt.Errorf("task completed at is required if completed true: %w", core_errors.ErrInvalidArgument)
		}
		if task.CompletedAt.Before(task.CreatedAt) {
			return fmt.Errorf("task completedAt must be greater than createdAt: %w", core_errors.ErrInvalidArgument)
		}
	} else {
		if task.CompletedAt != nil {
			return fmt.Errorf("task completedAt is required null if completed false: %w", core_errors.ErrInvalidArgument)
		}
	}
	if userId := task.UserID; userId < 0 {
		return fmt.Errorf("task user id must be greater than zero: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func NewTask(id int, version int, title string, description *string, completed bool, createdAt time.Time, completedAt *time.Time, userId int) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		UserID:      userId,
	}
}

func NewTaskUninitialized(title string, description *string, userID int) Task {
	return NewTask(UninitializedID, UninitializedVersion, title, description, false, time.Now(), nil, userID)
}
