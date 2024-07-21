package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iamalone98/todo-app/internal/http/helpers"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		tokenString := strings.Replace(token, "Bearer ", "", 1)

		claims, err := helpers.ParseJWTToken(tokenString)
		if err != nil {
			ctx.JSON(helpers.JSONErrorWrapper(http.StatusUnauthorized, "Auth required"))
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.Id)
		ctx.Set("login", claims.Login)

		ctx.Next()
	}
}
