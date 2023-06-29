package services

type ShopService struct {
	shopRepository repositories.ShopRepository
}

func NewShopService(shopRepository repositories.ShopRepository) *ShopService {
	return &ShopService{
		shopRepository: shopRepository,
	}
}

func (s *ShopService) GetAllShops() ([]entities.Shop, error) {
	// Implement the logic to fetch all shops from the repository
	// You can add any necessary validation or business logic here
	// Return the list of shops or an error if something goes wrong
	return s.shopRepository.GetAllShops()
}

func (s *ShopService) GetShopByID(id int) (*entities.Shop, error) {
	// Implement the logic to fetch a shop by ID from the repository
	// You can add any necessary validation or business logic here
	// Return the shop or an error if it doesn't exist or something goes wrong
	return s.shopRepository.GetShopByID(id)
}

func (s *ShopService) CreateShop(shop *entities.Shop) error {
	// Implement the logic to create a new shop in the repository
	// You can add any necessary validation or business logic here
	// Return an error if something goes wrong
	return s.shopRepository.CreateShop(shop)
}

func (s *ShopService) UpdateShop(shop *entities.Shop) error {
	// Implement the logic to update an existing shop in the repository
	// You can add any necessary validation or business logic here
	// Return an error if the shop doesn't exist or something goes wrong
	return s.shopRepository.UpdateShop(shop)
}

func (s *ShopService) DeleteShop(id int) error {
	// Implement the logic to delete a shop by ID from the repository
	// You can add any necessary validation or business logic here
	// Return an error if the shop doesn't exist or something goes wrong
	return s.shopRepository.DeleteShop(id)
}
