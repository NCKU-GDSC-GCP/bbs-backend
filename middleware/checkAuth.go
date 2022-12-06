package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CheckJWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := []byte(`{"Error":"Failed to process request, no token found"}`)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// TODO verify the JWT user given
		
	}
}