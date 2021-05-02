package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const DefaultPort = "8080"

func main() {
	// Create gin router
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Set up GET endpoint
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Golang Deployed to Heroku")
	})

	// Get port provided by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	// Start webapp
	log.Fatal(router.Run(":" + port))
}


