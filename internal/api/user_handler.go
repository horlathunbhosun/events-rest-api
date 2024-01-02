package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/models"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. Try again"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful", "data": user})
}
