package routes

import (
	"net/http"

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

	router.GET("/events/:id", getEvent)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
}
