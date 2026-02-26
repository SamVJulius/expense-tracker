package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              string           `gorm:"primaryKey"`
	UserID          string           `gorm:"not null;index"`
	CategoryID      string           `gorm:"not null;index"`
	Type            string         `gorm:"not null;index"` // income/expense
	Amount          float64        `gorm:"not null"`
	Note            string
	TransactionDate time.Time      `gorm:"not null;index"`
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
