package resolver

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
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
	return nil, nil
}

func (r *ShopResolver) ResolveGetShopByID(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r *ShopResolver) ResolveUpdateShop(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r *ShopResolver) ResolveDeleteShop(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
