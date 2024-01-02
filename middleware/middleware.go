package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/internal/utility"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {

		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token was found"})
		return
	}

	userId, err := utility.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unaurthorised"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
