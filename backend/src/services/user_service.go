package services

import (
	"errors"
	"time"
	"x-ui/backend/src/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := s.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) UpdateLastLogin(user *models.User) error {
	user.LastLogin = time.Now()
	return s.db.Save(user).Error
}

func (s *UserService) ChangePassword(userID uint, newPassword string) error {
	return s.db.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}