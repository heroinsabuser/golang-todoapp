package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/heroinsabuser/golang-todoapp/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key='%s' in path: %w", key, core_errors.ErrInvalidArgument)
	}

	value, err := strconv.Atoi(pathValue)

	if err != nil {
		return 0, fmt.Errorf("key='%s' pathValue='%s' not a valid integer in path: %v: %w", key, pathValue, err, core_errors.ErrInvalidArgument)
	}
	return value, nil
}
