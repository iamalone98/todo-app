package main

import (
	"log/slog"

	"github.com/iamalone98/todo-app/internal/db"
	"github.com/iamalone98/todo-app/internal/http/router"
	"github.com/iamalone98/todo-app/internal/http/server"
	"github.com/iamalone98/todo-app/internal/logger"
	"github.com/iamalone98/todo-app/internal/repository"
	"github.com/iamalone98/todo-app/internal/service"
	"github.com/lmittmann/tint"
)

func main() {
	logger.New()

	db, err := db.New()
	if err != nil {
		slog.Error("Database connect: ", tint.Err(err))
		return
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	r := router.New(service)
	server.New(r)
}
