package repositories

import (
	"github.com/justPoly/mini-wallet-service/backend/database"
	"github.com/justPoly/mini-wallet-service/backend/models"
)

func CreateAccount(account *models.Account) error {
	return database.DB.Create(account).Error
}

func GetAllAccounts() ([]models.Account, error) {

	var accounts []models.Account

	err := database.DB.Find(&accounts).Error

	return accounts, err
}

func GetAccountByID(id string) (*models.Account, error) {

	var account models.Account

	err := database.DB.First(&account, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func UpdateAccount(account *models.Account) error {
	return database.DB.Save(account).Error
}

func CreateTransaction(transaction *models.Transaction) error {
	return database.DB.Create(transaction).Error
}

func GetTransactionsByAccountID(accountID string) ([]models.Transaction, error) {

	var transactions []models.Transaction

	err := database.DB.
		Where("account_id = ?", accountID).
		Order("created_at DESC").
		Find(&transactions).Error

	return transactions, err
}