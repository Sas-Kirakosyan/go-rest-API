package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sas-Kirakosyan/my-go-project/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage": "could not fetch events, try again later"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "could not fetch event id."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	fmt.Printf("Event: %+v\n", event)
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse"})
	}

	event.ID = 1
	event.UserID = 1
	fmt.Printf("1111Event: %+v\n", event)
	event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage": "could not create event, try again later"})

		return
	}
}
