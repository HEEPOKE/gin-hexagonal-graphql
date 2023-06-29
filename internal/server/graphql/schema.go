package graphql

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/repositories"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/schemas"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
	"github.com/graphql-go/graphql"
)

var schema graphql.Schema

func init() {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)

	userResolver := resolver.NewUserResolver(userService)

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

	rootQuery := graphql.ObjectConfig{
		Name:   "Query",
		Fields: fields,
	}

	rootMutation := graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	var err error
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}
}

func GetSchema() *graphql.Schema {
	return &schema
}
