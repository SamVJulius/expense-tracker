package database

import "expense-tracker/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Transaction{},
	)
}
