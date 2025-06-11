package services

import (
	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: config.DB,
	}
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Create(user *models.User) error {
	return s.db.Create(user).Error
}
