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
	shops := []*models.Shop{}
	r.db.Find(&shops)
	return shops, nil
}

func (r *ShopRepository) CreateShop(shop *models.Shop) error {
	return nil
}

func (r *ShopRepository) GetShopByID(id int) (*models.Shop, error) {
	return nil, nil
}

func (r *ShopRepository) UpdateShop(shop *models.Shop) error {
	return nil
}

func (r *ShopRepository) DeleteShop(id int) error {
	return nil
}
