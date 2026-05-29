package users_transport_http

import (
	"net/http"

	core_logger "github.com/heroinsabuser/golang-todoapp/internal/core/logger"
	core_http_request "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/response"
)

type GetUserResponse UserResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get user id")
		return
	}
	userDomain, err := h.usersService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get user")
		return
	}
	response := GetUserResponse(userDtoFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
