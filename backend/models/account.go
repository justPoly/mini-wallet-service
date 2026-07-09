package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        string    `gorm:"type:text;primaryKey"`
	Name      string    `gorm:"not null"`
	Currency  string    `gorm:"size:3;not null"`
	Balance   float64   `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate runs automatically before a new account is inserted
func (a *Account) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New().String()
	return nil
}