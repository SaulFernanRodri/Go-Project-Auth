package services

import (
	"auth/models"
	"auth/repositories"
	"auth/utils"
	"errors"
)

type AuthServiceInterface interface {
	Login(email, password string) (string, error)
	Register(request models.RegisterRequest) error
}

type AuthService struct {
	repo *repositories.AuthRepo
}

func NewAuthService(repo *repositories.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user models.AuthUser

	// Buscar el usuario por username
	if err := s.repo.FindByUsername(username, &user); err != nil {
		return "", errors.New("user not found")
	}

	// Verificar la contraseña usando bcrypt
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generar un token JWT
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(request models.RegisterRequest) error {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := models.AuthUser{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	return s.repo.CreateUser(user)
}
