package handlers

import (
	"github.com/iamalone98/todo-app/internal/service"
)

type handlers struct {
	UserHandlers User
	TodoHandlers Todo
}

func NewHandlers(s service.Service) handlers {
	return handlers{
		UserHandlers: NewUserHandlers(s.UserService),
		TodoHandlers: NewTodoHandlers(s.TodoService),
	}
}
