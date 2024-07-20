package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamalone98/todo-app/internal/service"
)

type Todo interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	GetTodos(c *gin.Context)
	GetTodo(c *gin.Context)
}

type todo struct {
	s service.TodoService
}

func NewTodoHandlers(s service.TodoService) Todo {
	return todo{
		s: s,
	}
}

func (t todo) Create(ctx *gin.Context)   {}
func (t todo) Delete(ctx *gin.Context)   {}
func (t todo) Update(ctx *gin.Context)   {}
func (t todo) GetTodos(ctx *gin.Context) {}
func (t todo) GetTodo(ctx *gin.Context)  {}
