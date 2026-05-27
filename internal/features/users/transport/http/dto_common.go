package users_transport_http

import "github.com/heroinsabuser/golang-todoapp/internal/core/domain"

type UserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDtoFromDomain(user domain.User) UserResponse {
	return UserResponse{
		ID:          int(user.ID),
		Version:     int(user.Version),
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDtoFromDomains(users []domain.User) []UserResponse {
	res := make([]UserResponse, len(users))
	for i, user := range users {
		res[i] = userDtoFromDomain(user)
	}
	return res
}
