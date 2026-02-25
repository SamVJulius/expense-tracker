package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
	"expense-tracker/internal/repository"
	"expense-tracker/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

func Register(registerReq *dto.RegisterRequest) (string, error) {
	user, err := repository.GetByEmail(registerReq.Email)
	if err != nil {
		return "failed to fetch user", err
	}
	if user != nil {
		return "user already exists", nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return "failed to hash password", err
	}

	newUser := &models.User{
		Name:         registerReq.Name,
		Email:        registerReq.Email,
		PasswordHash: string(hash),
	}

	err = repository.Create(newUser)
	if err != nil {
		return "failed to create user", err
	}

	return "user registered successfully", nil
}

func Login(loginReq *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := repository.GetByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		return nil, err
	}

	tokenString, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: tokenString}, nil

}
