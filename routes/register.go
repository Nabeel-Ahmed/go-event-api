package routes

import (
	"net/http"
	"strconv"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered for event successfully."})

}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cancelled registration for event successfully."})
}
