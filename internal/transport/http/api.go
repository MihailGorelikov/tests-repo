package http

import (
	"log/slog"
	"net/http"
)

// New creates a new HTTP API.
func New(userService UserService, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users", CreateUserHandler(logger, userService))
	mux.HandleFunc("/api/v1/users/{id}", UsersHandler(logger, userService, "id"))

	return mux
}
