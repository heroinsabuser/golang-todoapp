package users_postgres_repository

import "github.com/heroinsabuser/golang-todoapp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func UserDomainsFromModels(models []UserModel) []domain.User {
	users := make([]domain.User, len(models))
	for i, model := range models {
		users[i] = domain.NewUser(model.ID, model.Version, model.FullName, model.PhoneNumber)
	}
	return users
}
