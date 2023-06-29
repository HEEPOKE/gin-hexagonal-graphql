package ports

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
	"github.com/graphql-go/graphql"
)

func ShopFields(shopResolver *resolver.ShopResolver) graphql.Fields {
	shopRepository := repositories.NewShopRepository(database.DB)

	fields := graphql.Fields{
		"shops": &graphql.Field{
			Type: graphql.NewList(schemas.ShopType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return shopResolver.ResolveGetAllShops(params)
			},
		},
		"shop": &graphql.Field{
			Type: schemas.ShopType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return shopResolver.ResolveGetShopByID(params)
			},
		},
		"createShop": &graphql.Field{
			Type: schemas.ShopType,
			Args: graphql.FieldConfigArgument{
				"shop": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.ShopInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return shopResolver.ResolveCreateShop(params)
			},
		},
		"updateShop": &graphql.Field{
			Type: schemas.ShopType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"shop": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.ShopInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return shopResolver.ResolveUpdateShop(params)
			},
		},
		"deleteShop": &graphql.Field{
			Type: schemas.ShopType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return shopResolver.ResolveDeleteShop(params)
			},
		},
	}

	return fields
}
