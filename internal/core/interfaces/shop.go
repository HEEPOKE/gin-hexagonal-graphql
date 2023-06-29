package interfaces

import "github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"

type ShopRepository interface {
	CreateShop(shop *models.Shop) error
	GetShopByID(id int) (*models.Shop, error)
	UpdateShop(shop *models.Shop) error
	DeleteShop(id int) error
}

type ShopService interface {
	CreateShop(shop *models.Shop) error
	GetShopByID(id int) (*models.Shop, error)
	UpdateShop(shop *models.Shop) error
	DeleteShop(id int) error
}
