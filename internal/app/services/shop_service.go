package services

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/repositories"
)

type ShopService struct {
	shopRepository *repositories.ShopRepository
}

func NewShopService(shopRepository *repositories.ShopRepository) *ShopService {
	return &ShopService{shopRepository: shopRepository}
}

func (s *ShopService) GetAllShops() ([]*models.Shop, error) {
	return s.shopRepository.GetAllShops()
}

func (s *ShopService) CreateShop(shop *models.Shop) (*models.Shop, error) {
	return s.shopRepository.CreateShop(shop)
}

func (s *ShopService) GetShopByID(id int) (*models.Shop, error) {
	return s.shopRepository.GetShopByID(id)
}

func (s *ShopService) UpdateShop(shop *models.Shop) (*models.Shop, error) {
	return s.shopRepository.UpdateShop(shop.ID, shop)
}

func (s *ShopService) DeleteShop(id int) error {
	return s.shopRepository.DeleteShop(id)
}
