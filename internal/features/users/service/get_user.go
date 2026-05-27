package users_service

import (
	"context"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
)

func (s *UsersService) GetUser(ctx context.Context, id int) (domain.User, error) {
	if id < 0 {
		return domain.User{}, core_errors.ErrInvalidArgument
	}

	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user failed: %w", err)
	}
	return user, nil
}
