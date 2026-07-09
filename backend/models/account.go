package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        string    `gorm:"type:text;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Currency  string    `gorm:"size:3;not null" json:"currency"`
	Balance   float64   `gorm:"default:0" json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New().String()
	return nil
}