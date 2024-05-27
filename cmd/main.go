package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/tests-repo/internal/platform/config"
	"github.com/tests-repo/internal/platform/server"
	"github.com/tests-repo/internal/repository"
	"github.com/tests-repo/internal/service"
	"github.com/tests-repo/internal/transport/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.ErrorContext(ctx, "failed to load config", slog.String("error", err.Error()))
		return
	}

	userRepository := repository.NewUser()
	userService := service.NewUser(userRepository)

	httpAPI := http.New(userService, logger)

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	server := server.New(addr, httpAPI)

	go func() {
		<-shutdownChan
		cancel()

		err = server.Shutdown(ctx)
		if err != nil {
			logger.ErrorContext(ctx, "failed to shutdown server", slog.String("error", err.Error()))
			return
		}
	}()

	logger.InfoContext(ctx, "starting server", slog.String("addr", addr))
	err = server.Start()
	if err != nil {
		logger.ErrorContext(ctx, "failed to start server", slog.String("error", err.Error()))
		return
	}
}
