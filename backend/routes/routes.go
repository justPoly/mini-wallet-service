package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/justPoly/mini-wallet-service/backend/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Mini Wallet API v2 - JSON TAG TEST",
		})
	})

	router.POST("/accounts", handlers.CreateAccount)
	router.GET("/accounts", handlers.GetAllAccounts)
	router.GET("/accounts/:id", handlers.GetAccount)
	router.POST("/accounts/:id/deposit", handlers.Deposit)

	router.POST("/transfers", handlers.Transfer)
	router.GET("/accounts/:id/transactions", handlers.GetTransactions)
}