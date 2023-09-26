// Package http contains HTTP related helper functions.
package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

// Server is a convenience wrapper around the standard
// library HTTP server.
type Server struct {
	srv *http.Server
}

// NewServer returns a server.
func NewServer(ctx context.Context, addr string, h http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
			Addr:              addr,
			Handler:           h,
			ReadHeaderTimeout: 2 * time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       120 * time.Second,
		},
	}
}

// Serve serves HTTP requests in a non-blocking way.
func (s *Server) Serve(errFn func(error)) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errFn(err)
		}
	}()
}

// Shutdown initiates the shutdown.
func (s *Server) Shutdown(timeout time.Duration) error {
	stopCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.srv.Shutdown(stopCtx)
}

// Close closes the server.
func (s *Server) Close() error {
	return s.srv.Close()
}
