package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamalone98/todo-app/internal/http/handlers"
	"github.com/iamalone98/todo-app/internal/service"
)

func New(s service.Service) *gin.Engine {
	h := handlers.NewHandlers(s)
	r := gin.Default()

	r.GET("/user", h.UserHandlers.Get)
	r.POST("/auth", h.UserHandlers.Authorization)
	r.POST("/reg", h.UserHandlers.Registration)

	return r
}
