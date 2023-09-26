package http

import (
	"net/http"

	"github.com/antiphp/gonew/internal/render"
)

// OKHandler returns an HTTP handler which always responds with HTTP 200 - OK.
func OKHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		render.JSON(rw, http.StatusOK, http.StatusText(http.StatusOK))
	})
}
