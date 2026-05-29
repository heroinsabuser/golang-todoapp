package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	// ASC direction by default here
	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	ORDER BY id
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel

		err := rows.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		userModels = append(userModels, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}
	userDomains := UserDomainsFromModels(userModels)
	return userDomains, nil
}
