package resolver

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/graphql-go/graphql"
)

type ShopResolver struct {
	shopService *services.ShopService
}

func NewShopResolver(shopService *services.ShopService) *ShopResolver {
	return &ShopResolver{shopService: shopService}
}

func (r *ShopResolver) ResolveGetAllShops(params graphql.ResolveParams) (interface{}, error) {
	return r.shopService.GetAllShops()
}

func (r *ShopResolver) ResolveCreateShop(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)
	description := params.Args["description"].(string)
	location := params.Args["location"].(string)
	telShop := params.Args["telShop"].(string)

	shop := &models.Shop{
		Name:        name,
		Description: description,
		Location:    location,
		TelShop:     telShop,
	}

	createdShop, err := r.shopService.CreateShop(shop)
	if err != nil {
		return nil, err
	}

	return createdShop, nil
}

func (r *ShopResolver) ResolveGetShopByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	shop, err := r.shopService.GetShopByID(id)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (r *ShopResolver) ResolveUpdateShop(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	name := params.Args["name"].(string)
	description := params.Args["description"].(string)
	location := params.Args["location"].(string)
	telShop := params.Args["telShop"].(string)

	shop := &models.Shop{
		ID:          id,
		Name:        name,
		Description: description,
		Location:    location,
		TelShop:     telShop,
	}

	updatedShop, err := r.shopService.UpdateShop(shop)
	if err != nil {
		return nil, err
	}

	return updatedShop, nil
}

func (r *ShopResolver) ResolveDeleteShop(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	err := r.shopService.DeleteShop(id)
	if err != nil {
		return nil, err
	}

	return true, nil
}
