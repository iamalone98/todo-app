package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamalone98/todo-app/internal/http/handlers"
	"github.com/iamalone98/todo-app/internal/http/middleware"
	"github.com/iamalone98/todo-app/internal/service"
)

func New(s service.Service) *gin.Engine {
	h := handlers.NewHandlers(s)
	r := gin.Default()

	r.GET("/user", h.UserHandlers.Get)
	r.POST("/auth", h.UserHandlers.Authorization)
	r.POST("/reg", h.UserHandlers.Registration)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.DELETE("/todo/:id", h.TodoHandlers.Delete)
		authorized.GET("/todo/:id", h.TodoHandlers.GetTodo)
		authorized.GET("/todos", h.TodoHandlers.GetTodos)
		authorized.POST("/todo", h.TodoHandlers.Create)
		authorized.PATCH("/todo", h.TodoHandlers.Update)
	}

	return r
}
