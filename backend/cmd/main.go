package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justPoly/mini-wallet-service/backend/database"
)

func main() {

	// Connect to the SQLite database
	database.ConnectDatabase()

	// Create the Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Mini Wallet API is running 🚀",
		})
	})

	// Start the server
	router.Run(":8080")
}