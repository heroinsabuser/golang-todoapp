package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/heroinsabuser/golang-todoapp/internal/core/logger"
	core_http_response "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/utils"
)

type GetUsersResponse []UserResponse

func (h *UsersHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid integer query parameter")
		return
	}
	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to fetch users")
		return
	}
	response := GetUsersResponse(usersDtoFromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	offset, err := core_http_utils.GetIntQueryParams(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("query param offset: %w", err)
	}
	limit, err := core_http_utils.GetIntQueryParams(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("query param limit: %w", err)
	}
	return limit, offset, nil
}
