package ports

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/schemas"
	"github.com/graphql-go/graphql"
)

func UserQueryFields(userResolver *resolver.UserResolver) graphql.Fields {
	fields := graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(schemas.UserType),
			Description: "Get all users",
			Resolve:     userResolver.ResolveGetAllUsers,
		},
		"user": &graphql.Field{
			Type:        schemas.UserType,
			Description: "Get a user by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: userResolver.ResolveGetUserByID,
		},
	}

	return fields
}

func UserMutationFields(userResolver *resolver.UserResolver) graphql.Fields {
	fields := graphql.Fields{
		"createUser": &graphql.Field{
			Type:        schemas.UserType,
			Description: "Create a new user",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.UserInputType),
				},
			},
			Resolve: userResolver.ResolveCreateUser,
		},
		"updateUser": &graphql.Field{
			Type:        schemas.UserType,
			Description: "Update an existing user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(schemas.UserInputType),
				},
			},
			Resolve: userResolver.ResolveUpdateUser,
		},
		"deleteUser": &graphql.Field{
			Type:        schemas.UserType,
			Description: "Delete a user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: userResolver.ResolveDeleteUser,
		},
		"getAllUsers": &graphql.Field{
			Type:        graphql.NewList(schemas.UserType),
			Description: "Get all users",
			Resolve:     userResolver.ResolveGetAllUsers,
		},
		"getUserByID": &graphql.Field{
			Type:        schemas.UserType,
			Description: "Get a user by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: userResolver.ResolveGetUserByID,
		},
	}

	return fields
}
