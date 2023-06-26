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

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepository.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(user *models.User) error {
	return s.UserRepository.DeleteUser(user)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.UserRepository.GetAllUsers()
}
