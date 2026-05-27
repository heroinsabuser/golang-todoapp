package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, version, full_name, phone_number
	FROM todoapp.users
	WHERE id = $1
	`
	row := r.pool.QueryRow(ctx, query, id)
	var userModel UserModel
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user by id='%d' in repository:%w", id, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan row: %w", err)
	}
	userDomain := domain.NewUser(userModel.ID, userModel.Version, userModel.FullName, userModel.PhoneNumber)
	return userDomain, nil
}
