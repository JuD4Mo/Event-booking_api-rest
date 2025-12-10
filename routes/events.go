package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	// Traer del contexto de Gin el userId
	userId := c.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	//Traer del contexto de Gin el userId
	userId := c.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event, the users do not match"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	//Traer del contexto de Gin el userId
	userId := c.GetInt64("userId")
	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event, the users do not match"})
		return
	}

	err = event.DeleteById()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
