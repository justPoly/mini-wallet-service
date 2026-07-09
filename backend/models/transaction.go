package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType string

const (
	DepositTransaction TransactionType = "DEPOSIT"
	TransferInTransaction TransactionType = "TRANSFER_IN"
	TransferOutTransaction TransactionType = "TRANSFER_OUT"
)

type Transaction struct {
	ID          string    `gorm:"type:text;primaryKey" json:"id"`
	AccountID   string    `gorm:"not null" json:"accountId"`
	Type        string    `gorm:"not null" json:"type"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	return nil
}