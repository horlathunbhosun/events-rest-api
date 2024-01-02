package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horlathunbhosun/events-rest-api/internal/api"
	"github.com/horlathunbhosun/events-rest-api/middleware"
)

func RegisterRoutes(server *gin.Engine) {

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "7704 page not found"})
	})

	server.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})
	//events routes
	v1 := server.Group("/v1")
	v1.GET("/events", api.GetEvents)
	v1.GET("/events/:eventId", api.GetSingleEventById)

	authedRoute := v1.Group("/")
	authedRoute.Use(middleware.Authenticate)
	authedRoute.POST("/events", api.CreateEvents)
	authedRoute.PUT("/events/:eventId", api.UpdateEvent)
	authedRoute.DELETE("/events/:eventId", api.DeleteEvent)

	authedRoute.POST("/events/:eventId/register", api.RegisterForEvent)
	authedRoute.DELETE("/events/:eventId/remove", api.CancelRegistration)

	//user routes
	v1.POST("/user/signup", api.Signup)
	v1.POST("/user/login", api.Login)
}
