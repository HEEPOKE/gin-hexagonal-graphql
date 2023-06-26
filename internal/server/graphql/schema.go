package graphql

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var createUserInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreateUserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var updateUserInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdateUserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Resolve logic for fetching all users
					userService := params.Context.Value("userService").(*services.UserService)
					return userService.GetAllUsers()
				},
			},
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Resolve logic for fetching a user by ID
					userService := params.Context.Value("userService").(*services.UserService)
					id := params.Args["id"].(int)
					return userService.GetUserByID(id)
				},
			},
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"user": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(createUserInputType),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Resolve logic for creating a user
					userService := params.Context.Value("userService").(*services.UserService)
					userInput := params.Args["user"].(map[string]interface{})
					name := userInput["name"].(string)
					email := userInput["email"].(string)

					user := &models.User{
						Username: name,
						Email:    email,
					}

					return userService.CreateUser(user)
				},
			},
			"updateUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"user": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(updateUserInputType),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Resolve logic for updating a user
					userService := params.Context.Value("userService").(*services.UserService)
					userInput := params.Args["user"].(map[string]interface{})
					id := userInput["id"].(int)
					name := userInput["name"].(string)
					email := userInput["email"].(string)

					user, err := userService.GetUserByID(id)
					if err != nil {
						return nil, err
					}

					user.Username = name
					user.Email = email

					return userService.UpdateUser(user)
				},
			},
			"deleteUser": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					userService := params.Context.Value("userService").(*services.UserService)
					id := params.Args["id"].(int)
					user, err := userService.GetUserByID(id)
					if err != nil {
						return nil, err
					}

					err = userService.DeleteUser(user)
					if err != nil {
						return nil, err
					}

					return true, nil
				},
			},
		},
	},
)

func NewSchema(userService *services.UserService) (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
