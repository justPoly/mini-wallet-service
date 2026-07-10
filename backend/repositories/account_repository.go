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

	err := database.DB.
		Order("created_at ASC").
		Find(&accounts).Error

	return accounts, err
}

func GetAccountByID(id string) (*models.Account, error) {

	var account models.Account

	err := database.DB.
		Where("id = ?", id).
		First(&account).Error

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

func GetTransactionsByAccountID(
	accountID string,
	limit int,
	offset int,
) ([]models.Transaction, error) {

	var transactions []models.Transaction

	err := database.DB.
		Where("account_id = ?", accountID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}