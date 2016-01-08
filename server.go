package main

import (
	"bitbucket.org/rking788/hotseats-api/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func statusHandler(c *gin.Context) {
	response := map[string]string{
		"msg":    "Let's Heat those Seats!!\n",
		"status": "Success",
	}
	c.JSON(http.StatusOK, response)
}

func main() {
	fmt.Println("Starting hotseats-api...")

	router := gin.New()

	// Configure middleware to be used
	router.Use(gin.Logger())

	// Route definitions
	router.GET("/status", statusHandler)
	router.GET("/events/:stadium", handlers.ListEvents)

	router.POST("/events", handlers.CreateEvent)

	port := ":8080"
	fmt.Println("Listening on port ", port)
	router.Run(port)
}
