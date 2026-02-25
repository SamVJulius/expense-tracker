package repository

import (
	"expense-tracker/internal/database"
	"expense-tracker/internal/models"
)

func Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetByID(id uint) (models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	return user, err
}