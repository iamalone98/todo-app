package service

import (
	"github.com/iamalone98/todo-app/internal/repository"
)

type Service struct {
	TodoService TodoService
	UserService UserService
}

func NewService(r repository.Repository) Service {
	userService := NewUserService(r.UserRepository)
	todoService := NewTodoService(r.TodoRepository)

	return Service{
		UserService: userService,
		TodoService: todoService,
	}
}
