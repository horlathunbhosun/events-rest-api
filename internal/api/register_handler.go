package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/models"
)

func RegisterForEvent(ctx *gin.Context) {
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

	err = eventData.Register(ctx.GetInt64("UserId"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration can not be completed."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registeration successful"})

}

func CancelRegistration(context *gin.Context) {

	eventId, _ := strconv.ParseInt(context.Param("eventId"), 10, 64)
	var event models.Event
	event.ID = eventId

	err := event.CancelRegistrationFromDb(context.GetInt64("UserId"))
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!!"})

}
