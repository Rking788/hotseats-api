package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rking788/hotseats-api/model"
	"net/http"
)

// ListEvents should return a JSON list of events for a particular stadium.
func ListEvents(c *gin.Context) {
	stadium := c.Param("stadium")

	if stadium == "" {
		errMsg := "Error: no stadium specified when requesting events!"
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, map[string]string{"status": errMsg})
	}

	fmt.Printf("Getting events for stadium: %s\n", stadium)

	evt := model.Event{EventType: "Foulball", Date: "2016-01-28T21:30:16GMT-0700"}
	evt2 := model.Event{EventType: "Homerun", Date: "2016-01-30T11:16:00GMT-0700"}

	eventList := make([]model.Event, 0, 10)
	eventList = append(eventList, evt)
	eventList = append(eventList, evt2)

	response := make(map[string][]model.Event)
	response["events"] = eventList

	c.JSON(http.StatusOK, response)
}

// CreateEvent is responsible for reading and parsing the JSON
// string from the request body and persisting it ... somewhere.
func CreateEvent(c *gin.Context) {
	var evt model.Event

	fmt.Printf("Content-Type from request: %s\n", c.ContentType())
	err := c.Bind(&evt)
	if err == nil {
		fmt.Printf("Creating event: %v...\n", evt)

		// TODO: Start actually persisting the event (evt) here...

		c.JSON(http.StatusCreated, map[string]string{"status": "Success"})
	} else {
		errMsg := fmt.Sprintf("Error: %s", err.Error())
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, map[string]string{"status": errMsg})
	}
}
