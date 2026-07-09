package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/justPoly/mini-wallet-service/backend/models"
	"github.com/justPoly/mini-wallet-service/backend/repositories"
)

type CreateAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Currency string `json:"currency" binding:"required,len=3"`
}

type DepositRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type TransferRequest struct {
	FromAccountID string  `json:"fromAccountId" binding:"required"`
	ToAccountID   string  `json:"toAccountId" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
}

func CreateAccount(c *gin.Context) {

	var request CreateAccountRequest

	// Read JSON request body
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	account := models.Account{
		Name:     request.Name,
		Currency: request.Currency,
	}

	err := repositories.CreateAccount(&account)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create account",
		})

		return
	}

	c.JSON(http.StatusCreated, account)
}

func GetAllAccounts(c *gin.Context) {

	accounts, err := repositories.GetAllAccounts()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch accounts",
		})

		return
	}

	c.JSON(http.StatusOK, accounts)
}

func GetAccount(c *gin.Context) {

	id := c.Param("id")

	account, err := repositories.GetAccountByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found",
		})

		return
	}

	c.JSON(http.StatusOK, account)
}

func Deposit(c *gin.Context) {

	id := c.Param("id")

	account, err := repositories.GetAccountByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found",
		})
		return
	}

	var request DepositRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	account.Balance += request.Amount

	err = repositories.UpdateAccount(account)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update balance",
		})
		return
	}

	transaction := models.Transaction{
		AccountID:   account.ID,
		Type:        models.DepositTransaction,
		Amount:      request.Amount,
		Description: "Account deposit",
	}

	err = repositories.CreateTransaction(&transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create transaction",
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func Transfer(c *gin.Context) {

	var request TransferRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if request.FromAccountID == request.ToAccountID {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot transfer to the same account",
		})

		return
	}

	fromAccount, err := repositories.GetAccountByID(request.FromAccountID)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sender account not found",
		})

		return
	}

	toAccount, err := repositories.GetAccountByID(request.ToAccountID)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Receiver account not found",
		})

		return
	}

	if fromAccount.Balance < request.Amount {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient balance",
		})

		return
	}

	fromAccount.Balance -= request.Amount

	toAccount.Balance += request.Amount

	if err := repositories.UpdateAccount(fromAccount); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update sender account",
		})

		return
	}

	if err := repositories.UpdateAccount(toAccount); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update receiver account",
		})

		return
	}

	outTransaction := models.Transaction{
		AccountID:   fromAccount.ID,
		Type:        models.TransferOutTransaction,
		Amount:      request.Amount,
		Description: "Transfer to " + toAccount.Name,
	}

	if err := repositories.CreateTransaction(&outTransaction); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create sender transaction",
		})

		return
	}

	inTransaction := models.Transaction{
		AccountID:   toAccount.ID,
		Type:        models.TransferInTransaction,
		Amount:      request.Amount,
		Description: "Transfer from " + fromAccount.Name,
	}

	if err := repositories.CreateTransaction(&inTransaction); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create receiver transaction",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer completed successfully",
		"from":    fromAccount,
		"to":      toAccount,
	})
}

func GetTransactions(c *gin.Context) {

	id := c.Param("id")

	_, err := repositories.GetAccountByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found",
		})

		return
	}

	transactions, err := repositories.GetTransactionsByAccountID(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve transactions",
		})

		return
	}

	c.JSON(http.StatusOK, transactions)
}