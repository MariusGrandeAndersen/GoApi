package service

import (
	"go_api/internal/model"
	"go_api/internal/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetAllUsers() ([]model.User, error) {
	return us.UserRepository.GetAllUsers()
}

func (us *UserService) CreateUser(user model.User) (model.User, error) {
	return us.UserRepository.CreateUser(user)
}
