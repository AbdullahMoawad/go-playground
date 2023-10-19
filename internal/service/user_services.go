package service

import (
	"app/internal/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("provided user is nil")
	}

	fmt.Println(user)
	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateByID(user *models.User, id string) error {
	if user == nil {
		return errors.New("provided user is nil")
	}
	if err := s.db.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
