package users_service

import (
	"context"
	"fmt"

	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
)

func (s *UsersService) DeleteUser(ctx context.Context, id int) error {
	if id < 0 {
		return core_errors.ErrInvalidArgument
	}

	err := s.usersRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("delete user failed: %w", err)
	}
	return nil
}
