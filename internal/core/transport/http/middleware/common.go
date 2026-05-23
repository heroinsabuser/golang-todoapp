package core_http_middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	core_logger "github.com/heroinsabuser/golang-todoapp/internal/core/logger"
	core_http_response "github.com/heroinsabuser/golang-todoapp/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const requestIdHeader = "X-Request-ID"


func RequestID() Middleware {
	return  func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdHeader)
			if requestId == ""{
				requestId = uuid.NewString()
			}

			r.Header.Set(requestIdHeader, requestId)
			w.Header().Set(requestIdHeader, requestId)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdHeader)

			l := log.With(
				zap.String("request_id", requestId),
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			log := core_logger.FromContext(ctx)

			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if err := recover(); err != nil {
					responseHandler.PanicResponse(err, "during handle HTTP request got unexpected panic")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}