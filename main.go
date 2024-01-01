package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/db"
	"github.com/horlathunbhosun/events-rest-api/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.Run(":9090")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvents(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
