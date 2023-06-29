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

func (s *ShopService) CreateShop(shop *models.Shop) error {
	return nil
}

func (s *ShopService) GetShopByID(id int) (*models.Shop, error) {
	return nil, nil
}

func (s *ShopService) UpdateShop(shop *models.Shop) error {
	return nil
}

func (s *ShopService) DeleteShop(id int) error {
	return nil
}
