package main

import (
	"log"
	"net/http"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		events := models.GetAllEvents()
		// Return JSON response
		c.JSON(http.StatusOK, events)
	})

	router.POST("/events", createEvent)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	event.ID = 1
	event.UserId = 1

	event.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
