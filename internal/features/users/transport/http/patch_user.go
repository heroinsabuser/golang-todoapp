package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/heroinsabuser/golang-todoapp/internal/core/domain"
	core_logger "github.com/heroinsabuser/golang-todoapp/internal/core/logger"
	core_http_request "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/response"
	core_http_types "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("full name is required")
		}
		fullNameLength := len([]rune(*r.FullName.Value))
		if fullNameLength < 3 || fullNameLength > 100 {
			return fmt.Errorf("full name must be 3 or more than 100 characters")
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLength := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLength < 10 || phoneNumberLength > 15 {
				return fmt.Errorf("phone number must be 10 or more than 15 characters")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("phone number must start with '+'")
			}
		}
	}
	return nil
}

type PatchUserResponse UserResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get user id")
		return
	}
	var req PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed get user request")
		return
	}
	userPatch := userPatchFromRequest(req)
	userDomain, err := h.usersService.PatchUser(ctx, userId, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed patch user")
		return
	}
	response := PatchUserResponse(userDtoFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(r PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(r.FullName.ToDomain(), r.PhoneNumber.ToDomain())
}
