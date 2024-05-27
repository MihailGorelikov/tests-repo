package http

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/tests-repo/internal/domain"
	"github.com/tests-repo/internal/transport/http/request"
	"github.com/tests-repo/internal/transport/http/response"
)

// UserService is a service for user.
type UserService interface {
	Create(ctx context.Context, email string) (domain.User, error)
	Get(ctx context.Context, id uuid.UUID) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func CreateUserHandler(logger *slog.Logger, svc UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request.CreateUser
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"invalid request body",
					slog.String("error", err.Error()),
				)

			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		user, err := svc.Create(r.Context(), req.Email)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"failed to create user",
					slog.String("email", req.Email),
					slog.String("error", err.Error()),
				)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := response.User{
			ID:    user.ID.String(),
			Email: user.Email,
		}

		res, err := json.Marshal(response)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"failed to marshal response",
					slog.String("error", err.Error()),
				)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func GetUserHandler(logger *slog.Logger, svc UserService, idKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(idKey)
		if id == "" {
			logger.
				ErrorContext(
					r.Context(),
					"missing id param",
					slog.String(idKey, id),
				)

			http.Error(w, "missing id param", http.StatusBadRequest)
			return
		}

		userID, err := uuid.Parse(id)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"invalid id param",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			http.Error(w, "invalid id param", http.StatusBadRequest)
			return
		}

		user, err := svc.Get(r.Context(), userID)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"failed to get user",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := response.User{
			ID:    user.ID.String(),
			Email: user.Email,
		}

		res, err := json.Marshal(response)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"failed to marshal response",
					slog.String("error", err.Error()),
				)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func DeleteUserHandler(logger *slog.Logger, svc UserService, idKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(idKey)
		if id == "" {
			logger.
				ErrorContext(
					r.Context(),
					"missing id param",
					slog.String(idKey, id),
				)

			http.Error(w, "missing id param", http.StatusBadRequest)
			return
		}

		userID, err := uuid.Parse(id)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"invalid id param",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			http.Error(w, "invalid id param", http.StatusBadRequest)
			return
		}

		err = svc.Delete(r.Context(), userID)
		if err != nil {
			logger.
				ErrorContext(
					r.Context(),
					"failed to delete user",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
