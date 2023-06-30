package repositories

import (
	"errors"

	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := ur.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := ur.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUser(id int, user *models.User) (*models.User, error) {
	err := ur.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return nil, err
	}

	updatedUser := &models.User{}
	err = ur.DB.First(updatedUser, id).Error
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (ur *UserRepository) DeleteUser(user *models.User) error {
	result := ur.DB.Delete(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}
