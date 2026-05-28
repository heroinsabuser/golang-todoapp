package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) PatchUser(ctx context.Context, id int, patch domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.users
	SET
	    full_name = $1,
	    phone_number = $2,
	    version = version + 1
	WHERE id = $3 AND version = $4
	RETURNING 
		id,
		version,
		full_name,
		phone_number
	`

	var userModel UserModel
	row := r.pool.QueryRow(ctx, query, patch.FullName, patch.PhoneNumber, id, patch.Version)
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id %d concurrently accessed: %w", id, core_errors.ErrConflict)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}
	userDomain := domain.NewUser(userModel.ID, userModel.Version, userModel.FullName, userModel.PhoneNumber)
	return userDomain, nil
}
