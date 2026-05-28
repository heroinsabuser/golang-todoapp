package users_service

import (
	"context"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
)

func (s *UsersService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	user, err := s.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed get user: %w", err)
	}
	err = user.ApplyPatch(patch)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed patch user: %w", err)
	}
	patchedUser, err := s.usersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed patch repository user: %w", err)
	}
	return patchedUser, nil
}
