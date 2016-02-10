package handlers

import (
	"fmt"
	"github.com/rking788/hotseats-api/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/rking788/hotseats-api/db"
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

	// Store the new event in the DB
	dbConn, dbErr := db.GetDBConnection()
	if dbErr != nil {
		fmt.Printf("Found err: %s\n", dbErr.Error())
		c.JSON(http.StatusInternalServerError,
			map[string]string{"status": "Error cannot connect to DB!"})
		return
	}

	// Find the `sid` by which the events will be selected
	var outStadium model.Stadium
	dbConn.Table("stadiums").Where("name = ?", stadium).First(&outStadium)

	fmt.Printf("Found sid=%d\n", outStadium.Sid)
	fmt.Printf("Getting events for stadium: %s\n", stadium)

	// Select all events in the events table by the `sid`
	eventList := make([]model.Event, 0, 3)
	dbConn.Where(&model.Event{Sid: outStadium.Sid}).Find(&eventList)

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

		// Store the new event in the DB
		dbConn, dbErr := db.GetDBConnection()
		if dbErr != nil {
			fmt.Printf("Found err: %s\n", dbErr.Error())
			c.JSON(http.StatusInternalServerError,
				map[string]string{"status": "Error cannot connect to DB!"})
			return
		}

		// Find the `sid` value that should be used when inserting the new event
		var outStadium model.Stadium
		dbConn.Table("stadiums").Where("name = ?", evt.Stadium.Name).First(&outStadium)

		fmt.Printf("Found sid=%d\n", outStadium.Sid)

		if outStadium.Sid == 0 {
			fmt.Println("Didn't find a stadium with that name!!")
			c.JSON(http.StatusBadRequest, map[string]string{"status": "Error no stadium with that name"})
			return
		}

		evt.Sid = outStadium.Sid
		dbConn.Create(&evt)

		c.JSON(http.StatusCreated, map[string]string{"status": "Success"})
	} else {
		errMsg := fmt.Sprintf("Error: %s", err.Error())
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, map[string]string{"status": errMsg})
	}
}
