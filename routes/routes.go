package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/internal/api"
)

func RegisterRoutes(server *gin.Engine) {

	//events routes
	server.GET("/events", api.GetEvents)
	server.GET("/events/:eventId", api.GetSingleEventById)
	server.POST("/events", api.CreateEvents)
	server.PUT("/events/:eventId", api.UpdateEvent)
	server.DELETE("/events/:eventId", api.DeleteEvent)

	//user routes

	server.POST("/user/signup")
}
