package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamalone98/todo-app/internal/http/helpers"
	"github.com/iamalone98/todo-app/internal/service"
	"github.com/iamalone98/todo-app/models"
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

func (t todo) Create(ctx *gin.Context) {
	var todo models.TodoCreate

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Missing required params"))
		return
	}

	userId, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Unknown error"))
		return
	}

	todoRet, err := t.s.Create(userId.(int), todo)
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusNotFound, "Failed create todo"))
		return
	}

	ctx.JSON(http.StatusOK, todoRet)
}
func (t todo) Delete(ctx *gin.Context) {
	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)

	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Invalid todo ID"))
		return
	}

	userId, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Unknown error"))
		return
	}

	todo, err := t.s.Delete(todoId, userId.(int))
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusNotFound, "Todo not found"))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (t todo) Update(ctx *gin.Context) {
	var todo models.TodoUpdate

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Missing required params"))
		return
	}

	userId, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Unknown error"))
		return
	}

	todoRet, err := t.s.Update(userId.(int), todo)
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusNotFound, "Failed update todo"))
		return
	}

	ctx.JSON(http.StatusOK, todoRet)
}

func (t todo) GetTodos(ctx *gin.Context) {
	userId, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Unknown error"))
		return
	}

	todos, err := t.s.GetAll(userId.(int))
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusNotFound, "Failed get todos"))
		return
	}

	ctx.JSON(http.StatusOK, todos)
}
func (t todo) GetTodo(ctx *gin.Context) {
	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)

	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Invalid todo ID"))
		return
	}

	userId, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Unknown error"))
		return
	}

	todoRet, err := t.s.Get(todoId, userId.(int))
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusNotFound, "Todo not found"))
		return
	}

	ctx.JSON(http.StatusOK, todoRet)
}
