package ports

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/schemas"
	"github.com/graphql-go/graphql"
)

func ShopQueryFields(shopResolver *resolver.ShopResolver) graphql.Fields {
	fields := graphql.Fields{
		"shops": &graphql.Field{
			Type:        graphql.NewList(schemas.ShopType),
			Description: "Get all shops",
			Resolve:     shopResolver.ResolveGetAllShops,
		},
		"shop": &graphql.Field{
			Type:        schemas.ShopType,
			Description: "Get a shop by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: shopResolver.ResolveGetShopByID,
		},
	}

	return fields
}

func ShopMutationFields(shopResolver *resolver.ShopResolver) graphql.Fields {
	fields := graphql.Fields{
		"createShop": &graphql.Field{
			Type:        schemas.ShopType,
			Description: "Create a new shop",
			Args: graphql.FieldConfigArgument{
				"shop": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.ShopInputType),
				},
			},
			Resolve: shopResolver.ResolveCreateShop,
		},
		"updateShop": &graphql.Field{
			Type:        schemas.ShopType,
			Description: "Update an existing shop",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"shop": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.ShopInputType),
				},
			},
			Resolve: shopResolver.ResolveUpdateShop,
		},
		"deleteShop": &graphql.Field{
			Type:        schemas.ShopType,
			Description: "Delete a shop",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: shopResolver.ResolveDeleteShop,
		},
	}

	return fields
}
