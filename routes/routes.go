package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // events/1, events/5
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent) // events/1, events/5
	server.DELETE("events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
