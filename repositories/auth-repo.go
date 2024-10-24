package repositories

import (
	"auth/models"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	FindByUsername(username string, user *models.AuthUser) error
	CreateUser(user models.AuthUser) error
}

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (repo *AuthRepo) FindByUsername(username string, user *models.AuthUser) error {
	result := repo.db.Where("username = ?", username).First(user)
	return result.Error
}

func (repo *AuthRepo) CreateUser(user models.AuthUser) error {
	return repo.db.Create(&user).Error
}
