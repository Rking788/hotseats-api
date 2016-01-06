package handlers

import (
	"bitbucket.org/rking788/hotseats-api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListEvents should return a JSON list of events for a particular stadium.
func ListEvents(c *gin.Context) {
	stadium := c.Param("stadium")
	fmt.Printf("Getting events for stadium: %s\n", stadium)

	evt := model.Event{EventType: "Foulball", Date: "2016-01-31"}
	evt2 := model.Event{EventType: "Homerun", Date: "2016-01-30"}

	eventList := make([]model.Event, 0, 10)
	eventList = append(eventList, evt)
	eventList = append(eventList, evt2)

	response := make(map[string][]model.Event)
	response["events"] = eventList

	c.JSON(http.StatusOK, response)
}
