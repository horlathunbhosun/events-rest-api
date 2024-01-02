package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/internal/utility"
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

func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong request invalid"})
		return
	}

	err = user.ValidateUserCredential()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utility.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messsage": "Login successful", "data": token})
}
