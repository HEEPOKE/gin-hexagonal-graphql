package resolver

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models/response"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/constants"
	"github.com/graphql-go/graphql"
)

type ShopResolver struct {
	shopService *services.ShopService
}

func NewShopResolver(shopService *services.ShopService) *ShopResolver {
	return &ShopResolver{shopService: shopService}
}

func (r *ShopResolver) ResolveGetAllShops(params graphql.ResolveParams) (interface{}, error) {
	shops, err := r.shopService.GetAllShops()
	if err != nil {
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(shops, constants.SUCCESS)
	return successResponse, nil
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
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(createdShop, constants.SUCCESS)
	return successResponse, nil
}

func (r *ShopResolver) ResolveGetShopByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	shop, err := r.shopService.GetShopByID(id)
	if err != nil {
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(shop, constants.SUCCESS)
	return successResponse, nil
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
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(updatedShop, constants.SUCCESS)
	return successResponse, nil
}

func (r *ShopResolver) ResolveDeleteShop(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	err := r.shopService.DeleteShop(id)
	if err != nil {
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(true, constants.SUCCESS)
	return successResponse, nil
}
