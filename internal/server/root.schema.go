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
	userFields := ports.UserFields(userResolver)

	shopRepository := repositories.NewShopRepository(database.DB)
	shopService := services.NewShopService(shopRepository)
	shopResolver := resolver.NewShopResolver(shopService)
	shopFields := ports.ShopFields(shopResolver)

	fields := graphql.Fields{
		"users":      userFields["users"],
		"user":       userFields["user"],
		"createUser": userFields["createUser"],
		"updateUser": userFields["updateUser"],
		"deleteUser": userFields["deleteUser"],
		"shops":      shopFields["shops"],
		"shop":       shopFields["shop"],
		"createShop": shopFields["createShop"],
		"updateShop": shopFields["updateShop"],
		"deleteShop": shopFields["deleteShop"],
	}

	return fields
}
