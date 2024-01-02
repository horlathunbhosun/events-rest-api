package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/models"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data. Try again"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": events})
}

func GetSingleEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventId"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch save data. Try again"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": event})

}

func CreateEvents(ctx *gin.Context) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
		return
	}
	event.UserId = ctx.GetInt64("userId")

	err = event.Save()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event. Try again"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "data": event})
}

func UpdateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventId"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId"})
		return
	}
	eventData, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with that id."})
		return
	}

	if eventData.UserId != ctx.GetInt64("UserId") {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You can not update this event"})
		return
	}

	var updateEvent models.Event
	err = ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong, Bad Request"})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully.", "data": updateEvent})
}

func DeleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventId"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId"})
		return
	}
	eventData, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event id does not exist."})
		return
	}

	if eventData.UserId != ctx.GetInt64("UserId") {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You can not delete this event"})
		return
	}

	err = eventData.Delete()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event delete successfully."})

}
