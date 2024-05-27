package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New creates a new HTTP API.
func New(userService UserService, logger *slog.Logger) http.Handler {
	engine := gin.Default()

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.POST("/", CreateUser(logger, userService))
				users.GET("/:id", GetUser(logger, userService, "id"))
				users.DELETE("/:id", DeleteUser(logger, userService, "id"))
			}
		}
	}

	return engine
}
