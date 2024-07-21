package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iamalone98/todo-app/internal/http/helpers"
	"github.com/iamalone98/todo-app/internal/service"
	"github.com/iamalone98/todo-app/internal/utils"
	"github.com/iamalone98/todo-app/models"
)

type User interface {
	Authorization(c *gin.Context)
	Registration(c *gin.Context)
	Get(c *gin.Context)
}

type user struct {
	s service.UserService
}

func NewUserHandlers(s service.UserService) User {
	return user{
		s: s,
	}
}

func (u user) Authorization(ctx *gin.Context) {
	var user models.UserAuth

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Missing required params"))
		return
	}

	userData, err := u.s.Get(*user.Login)
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusNotFound, "User not found"))
		return
	}

	if err := utils.CheckHashPassword([]byte(*userData.Password), []byte(*user.Password)); err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusUnauthorized, "Wrong password"))
		return
	}

	token, err := helpers.GenerateJWTToken(*userData.Id, *user.Login)
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Error create token"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u user) Registration(ctx *gin.Context) {
	var user models.UserAuth

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Missing required params"))
		return
	}

	if len(*user.Password) < 8 {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusBadRequest, "Min password length 8"))
		return
	}

	newUser, err := u.s.Create(user)
	if err != nil && strings.Contains(err.Error(), "duplicate key value") {
		if strings.Contains(err.Error(), "duplicate key value") {
			ctx.JSON(helpers.JSONErrorWrapper(http.StatusConflict, "This login is already taken"))
			return
		}

		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Unknown error"))
	}

	token, err := helpers.GenerateJWTToken(*newUser.Id, *newUser.Login)
	if err != nil {
		ctx.JSON(helpers.JSONErrorWrapper(http.StatusInternalServerError, "Error create token"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u user) Get(ctx *gin.Context) {}
