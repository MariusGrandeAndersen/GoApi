package repository

import (
	"go_api/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := ur.DB.Find(&users).Error
	return users, err
}

func (ur *UserRepository) CreateUser(user model.User) (model.User, error) {
	err := ur.DB.Create(&user).Error
	return user, err
}
