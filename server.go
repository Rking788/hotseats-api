package main

import (
	"bitbucket.org/rking788/hotseats-api/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func statusHandler(c *gin.Context) {
	c.String(http.StatusOK, "Let's Heat those Seats!!\n")
}

func main() {
	fmt.Println("Hello, World!")

	router := gin.New()

	// Configure middleware to be used
	router.Use(gin.Logger())

	// Route definitions
	router.GET("/status", statusHandler)
	router.GET("/events/:stadium", handlers.ListEvents)

	port := ":8080"
	fmt.Println("Listening on port ", port)
	router.Run(port)
}
