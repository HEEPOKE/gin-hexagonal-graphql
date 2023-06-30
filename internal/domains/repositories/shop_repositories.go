package repositories

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"gorm.io/gorm"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{db: db}
}

func (r *ShopRepository) GetAllShops() ([]*models.Shop, error) {
	var shops []*models.Shop
	result := r.db.Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}
	return shops, nil
}

func (r *ShopRepository) GetShopByID(id int) (*models.Shop, error) {
	shop := &models.Shop{}
	result := r.db.First(shop, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return shop, nil
}

func (r *ShopRepository) CreateShop(shop *models.Shop) (*models.Shop, error) {
	result := r.db.Create(shop)
	if result.Error != nil {
		return nil, result.Error
	}
	return shop, nil
}

func (r *ShopRepository) UpdateShop(id int, shop *models.Shop) (*models.Shop, error) {
	err := r.db.Model(&models.Shop{}).Where("id = ?", id).Updates(shop).Error
	if err != nil {
		return nil, err
	}

	updatedShop := &models.Shop{}
	err = r.db.First(updatedShop, id).Error
	if err != nil {
		return nil, err
	}

	return updatedShop, nil
}

func (r *ShopRepository) DeleteShop(id int) error {
	shop := &models.Shop{}
	result := r.db.First(shop, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}
	result = r.db.Delete(shop)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
