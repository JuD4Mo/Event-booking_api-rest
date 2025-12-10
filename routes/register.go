package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
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

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messsage": "could not register user for event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messsage": "could not cancel registration for event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "canceled"})
}
