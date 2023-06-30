package repositories

import (
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
	return users, result.Error
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	result := ur.DB.Create(user)
	return result.Error
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := ur.DB.First(&user, id)
	return &user, result.Error
}

func (ur *UserRepository) UpdateUser(id int, user *models.User) error {
	err := ur.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
	return err
}

func (ur *UserRepository) DeleteUser(user *models.User) error {
	result := ur.DB.Delete(user)
	return result.Error
}
