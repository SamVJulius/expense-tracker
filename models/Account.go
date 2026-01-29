package models

import "time"

type Account struct {
	AccountID string `gorm:"primaryKey" json:"accountid"`

	UserID string `gorm:"not null;uniqueIndex" json:"userid"`
	User   *User   `gorm:"foreignKey:UserID"`

	Balance int `json:"balance"`
	Budget  int `json:"budget"`

	Transactions []Transaction `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transactions"`

	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
