// Package api contains the HTTP server.
package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/antiphp/gonew/internal/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hamba/logger/v2"
)

// Foobar represents the foobar application.
type Foobar interface {
	Foobar(ctx context.Context, msg string) (string, error)
}

// Server serves HTTP requests.
type Server struct {
	h http.Handler

	app Foobar

	log *logger.Logger
}

// NewServer returns a new HTTP server.
func NewServer(app Foobar, log *logger.Logger) *Server {
	srv := &Server{
		app: app,
		log: log,
	}

	srv.h = srv.routes()

	return srv
}

// ServeHTTP serves http requests.
func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.h.ServeHTTP(rw, r)
}

func (s *Server) routes() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.RequestID)

	mux.Post("/foobar", s.handleFoobar())

	mux.NotFound(handleJSON(http.StatusNotFound))
	mux.MethodNotAllowed(handleJSON(http.StatusMethodNotAllowed))
	return mux
}

func handleJSON(status int) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		render.JSONError(rw, status, strings.ToLower(http.StatusText(status)))
	}
}
