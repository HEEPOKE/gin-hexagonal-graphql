package ports

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/schemas"
	"github.com/graphql-go/graphql"
)

func UserFields(userResolver *resolver.UserResolver) graphql.Fields {
	fields := graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(schemas.UserType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return userResolver.ResolveGetAllUsers(params)
			},
		},
		"user": &graphql.Field{
			Type: schemas.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return userResolver.ResolveGetUserByID(params)
			},
		},
		"createUser": &graphql.Field{
			Type: schemas.UserType,
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.UserInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return userResolver.ResolveCreateUser(params)
			},
		},
		"updateUser": &graphql.Field{
			Type: schemas.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.UserInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return userResolver.ResolveUpdateUser(params)
			},
		},
		"deleteUser": &graphql.Field{
			Type: schemas.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return userResolver.ResolveDeleteUser(params)
			},
		},
	}

	return fields
}
