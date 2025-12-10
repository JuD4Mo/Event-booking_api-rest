package routes

import (
	"net/http"

	"example.com/event-booking/middlewares"
	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define a simple GET endpoint
	router.GET("/events", func(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
			return
		}
		// Return JSON response
		c.JSON(http.StatusOK, events)
	})

	// Ejecutar el middleware antes de la operación como tal:

	//Crear un grupo de peticiones
	authenticated := router.Group("/")

	//Ese grupo usa el middleware
	authenticated.Use(middlewares.Authenticate)

	//Se generan las peticiones en ese grupo
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)

	router.GET("/events/:id", getEvent)
	router.POST("/signup", signup)
	router.POST("/login", login)
}
