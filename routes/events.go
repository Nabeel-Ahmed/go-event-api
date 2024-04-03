package routes

import (
	"net/http"
	"strconv"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}

	context.JSON(200, events)
}

func createEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully.", "event": event})
}

func getEventById(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID.", "message": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID.", "message": err.Error()})
		return
	}

	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event."})
		return
	}

	updatedEvent := models.Event{}
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data.", "message": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	if err := updatedEvent.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event.", "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})

}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID.", "message": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event.", "message": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event.", "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
