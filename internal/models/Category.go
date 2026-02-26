package models

import "time"

type Category struct {
	ID        string      `gorm:"primaryKey"`
	UserID    string      `gorm:"not null;uniqueIndex:idx_user_category"`
	Name      string    `gorm:"not null;uniqueIndex:idx_user_category"`
	Type      string    `gorm:"not null;index"` 
	CreatedAt time.Time
}
