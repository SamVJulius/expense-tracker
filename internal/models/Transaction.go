package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint           `gorm:"primaryKey"`
	UserID          uint           `gorm:"not null;index"`
	CategoryID      uint           `gorm:"not null;index"`
	Type            string         `gorm:"not null;index"` // income/expense
	Amount          float64        `gorm:"not null"`
	Note            string
	TransactionDate time.Time      `gorm:"not null;index"`
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
