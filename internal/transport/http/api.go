package http

import (
	"log/slog"
	"net/http"
)

// New creates a new HTTP API.
func New(userService UserService, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users", UsersHandler(logger, userService))

	return mux
}
