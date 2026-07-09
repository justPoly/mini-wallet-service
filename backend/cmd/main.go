package main

import (
	"github.com/gin-gonic/gin"

	"github.com/justPoly/mini-wallet-service/backend/database"
	"github.com/justPoly/mini-wallet-service/backend/routes"
)

func main() {

	// Connect database
	database.ConnectDatabase()

	// Create router
	router := gin.Default()

	// Register application routes
	routes.RegisterRoutes(router)

	// Start server
	router.Run(":8080")
}