package server

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/core/ports"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/repositories"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	userResolver := resolver.NewUserResolver(userService)
	userQueryFields := ports.UserQueryFields(userResolver)

	shopRepository := repositories.NewShopRepository(database.DB)
	shopService := services.NewShopService(shopRepository)
	shopResolver := resolver.NewShopResolver(shopService)
	shopQueryFields := ports.ShopQueryFields(shopResolver)

	fields := graphql.Fields{
		"users": userQueryFields["users"],
		"user":  userQueryFields["user"],
		"shops": shopQueryFields["shops"],
		"shop":  shopQueryFields["shop"],
	}

	return fields
}

func GetRootMutationFields() graphql.Fields {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	userResolver := resolver.NewUserResolver(userService)
	userMutationFields := ports.UserMutationFields(userResolver)

	shopRepository := repositories.NewShopRepository(database.DB)
	shopService := services.NewShopService(shopRepository)
	shopResolver := resolver.NewShopResolver(shopService)
	shopMutationFields := ports.ShopMutationFields(shopResolver)

	fields := graphql.Fields{
		"createUser": userMutationFields["createUser"],
		"updateUser": userMutationFields["updateUser"],
		"deleteUser": userMutationFields["deleteUser"],
		"createShop": shopMutationFields["createShop"],
		"updateShop": shopMutationFields["updateShop"],
		"deleteShop": shopMutationFields["deleteShop"],
	}

	return fields
}
