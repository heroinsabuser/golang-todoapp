package users_transport_http

import (
	"net/http"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
	core_logger "github.com/heroinsabuser/golang-todoapp/internal/core/logger"
	core_http_request "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type CreateUserResponse UserResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	var req CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate create user request")
		return
	}
	userDomain := domainFromDto(req)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(userDtoFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDto(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
