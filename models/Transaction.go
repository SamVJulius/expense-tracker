package models

import "time"

type Transaction struct {
	TransactionID string `gorm:"primaryKey" json:"transactionid"`

	AccountID string  `gorm:"not null;index" json:"accountid"`
	Account   *Account `gorm:"foreignKey:AccountID"`

	Type        string `json:"type"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`

	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
