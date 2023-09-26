// Package render renders HTTP responses.
package render

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// Error represents a JSON response error.
type Error struct {
	Code    int    `json:"errno"`
	Message string `json:"error"`
}

// JSON renders a JSON response.
func JSON(rw http.ResponseWriter, code int, val any) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)

	if err := jsoniter.NewEncoder(rw).Encode(val); err != nil {
		_, _ = rw.Write([]byte(`{"errno":500,"error":"internal server error"}`))
	}
}

// JSONError renders a JSON error response.
func JSONError(rw http.ResponseWriter, code int, status string) {
	JSON(rw, code, Error{
		Code:    code,
		Message: status,
	})
}
