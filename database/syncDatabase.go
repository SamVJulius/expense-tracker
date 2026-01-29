package database

import "expense-tracker/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
	)
}
