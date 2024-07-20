package service

import "github.com/iamalone98/todo-app/internal/repository"

type Service struct {
	TodoService TodoService
	UserService UserService
}

func NewService(r repository.Repository) Service {
	return Service{
		TodoService: NewTodoService(r.TodoRepository),
		UserService: NewUserService(r.UserRepository),
	}
}
