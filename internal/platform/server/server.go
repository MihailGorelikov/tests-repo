package server

import (
	"context"
	"net/http"
	"time"
)

// Server is a HTTP server.
type Server struct {
	srv *http.Server
}

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

// New creates a new HTTP server.
func New(addr string, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}
}

// Start starts the server.
func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
