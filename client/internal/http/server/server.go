package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lmittmann/tint"
)

func New(r http.Handler) {
	addr := ":3000"
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		slog.Info("Starting server", "addr:", addr)

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server:", tint.Err(err))
		}

		slog.Info("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("HTTP shutdown:", tint.Err(err))
	}

	slog.Info("Graceful shutdown complete.")
}
