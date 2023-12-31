package services

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetAllUsers() ([]*models.User, error) {
	return us.UserRepository.GetAllUsers()
}

func (us *UserService) CreateUser(user *models.User) error {
	err := us.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByID(id int) (*models.User, error) {
	return us.UserRepository.GetUserByID(id)
}

func (us *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return us.UserRepository.UpdateUser(user.ID, user)
}

func (us *UserService) DeleteUser(user *models.User) error {
	return us.UserRepository.DeleteUser(user)
}
