package http

import (
	"log/slog"
	"net/http"
)

// New creates a new HTTP API.
func New(userService UserService, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users", CreateUserHandler(logger, userService))
	mux.HandleFunc("GET /api/v1/users/{id}", GetUserHandler(logger, userService, "id"))
	mux.HandleFunc("DELETE /api/v1/users/{id}", DeleteUserHandler(logger, userService, "id"))

	return mux
}
