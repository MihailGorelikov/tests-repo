package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
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

// GetUser gets a user.
func GetUser(logger *slog.Logger, svc UserService, idKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(idKey)
		if id == "" {
			logger.
				ErrorContext(
					c,
					"missing id param",
					slog.String("id", id),
				)

			c.JSON(http.StatusBadRequest, response.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "missing id param",
			})
			return
		}

		userID, err := uuid.Parse(id)
		if err != nil {
			logger.
				ErrorContext(
					c,
					"invalid id param",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			c.JSON(http.StatusBadRequest, response.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "invalid id param",
			})
			return
		}

		user, err := svc.Get(c, userID)
		if err != nil {
			logger.
				ErrorContext(
					c,
					"failed to get user",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			c.JSON(http.StatusInternalServerError, response.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response.User{
			ID:    user.ID.String(),
			Email: user.Email,
		})
	}
}

// CreateUser creates a user.
func CreateUser(logger *slog.Logger, svc UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateUser
		if err := c.BindJSON(&req); err != nil {
			logger.
				ErrorContext(
					c,
					"invalid request body",
					slog.String("error", err.Error()),
				)

			c.JSON(http.StatusBadRequest, response.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "invalid request body",
			})
			return
		}

		user, err := svc.Create(c, req.Email)
		if err != nil {
			logger.
				ErrorContext(
					c,
					"failed to create user",
					slog.String("email", req.Email),
					slog.String("error", err.Error()),
				)

			c.JSON(http.StatusInternalServerError, response.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, response.User{
			ID:    user.ID.String(),
			Email: user.Email,
		})
	}
}

// DeleteUser deletes a user.
func DeleteUser(logger *slog.Logger, svc UserService, idKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(idKey)
		if id == "" {
			logger.
				ErrorContext(
					c,
					"missing id param",
					slog.String("id", id),
				)

			c.JSON(http.StatusBadRequest, response.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "missing id param",
			})
			return
		}

		userID, err := uuid.Parse(id)
		if err != nil {
			logger.
				ErrorContext(
					c,
					"invalid id param",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			c.JSON(http.StatusBadRequest, response.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "invalid id param",
			})
			return
		}

		err = svc.Delete(c, userID)
		if err != nil {
			logger.
				ErrorContext(
					c,
					"failed to delete user",
					slog.String("id", id),
					slog.String("error", err.Error()),
				)

			c.JSON(http.StatusInternalServerError, response.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}

		c.Status(http.StatusOK)
	}
}
