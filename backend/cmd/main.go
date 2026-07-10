package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/justPoly/mini-wallet-service/backend/database"
	"github.com/justPoly/mini-wallet-service/backend/routes"
)

func main() {

	database.ConnectDatabase()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Accept",
			"Authorization",
			"Idempotency-Key",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}