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
	ID          string          `gorm:"type:text;primaryKey"`
	AccountID   string          `gorm:"not null;index"`
	Type        TransactionType `gorm:"not null"`
	Amount      float64         `gorm:"not null"`
	Description string
	CreatedAt   time.Time
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	return nil
}