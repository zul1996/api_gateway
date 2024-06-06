package model

import (
	"time"
)

type Transaction struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	AccountID       string    `gorm:"not null" json:"account_id"`
	BankID          string    `gorm:"not null" json:"bank_id"`
	Amount          float64   `gorm:"not null" json:"amount"`
	TransactionDate time.Time `gorm:"type:timestamp with time zone;not null" json:"transaction_date"`
}

func (Transaction) TableName() string {
	return "transaction"
}
