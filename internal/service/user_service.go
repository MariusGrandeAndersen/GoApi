package service

import (
	"go_api/internal/model"
	"go_api/internal/repository"
	"strings"
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

// ConvertNameToUpperCase converts the given name to uppercase - useless but used in test coverage
func ConvertNameToUpperCase(name string) string {
	return strings.ToUpper(name)
}
