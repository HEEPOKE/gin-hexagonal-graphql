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
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"tel": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var userInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"username": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"tel": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"role": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
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
					userService := params.Context.Value("userService").(*services.UserService)
					id := int(params.Args["id"].(int))
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
						Type: graphql.NewNonNull(userInputType),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					userService := params.Context.Value("userService").(*services.UserService)
					userInput := params.Args["user"].(map[string]interface{})
					email := userInput["email"].(string)
					username := userInput["username"].(string)
					password := userInput["password"].(string)
					tel, _ := userInput["tel"].(string)
					role, _ := userInput["role"].(string)

					user := &models.User{
						Email:    email,
						Username: username,
						Password: password,
						Tel:      tel,
						Role:     role,
					}

					err := userService.CreateUser(user)
					if err != nil {
						return nil, err
					}

					return user, nil
				},
			},
			"updateUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"user": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(userInputType),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					userService := params.Context.Value("userService").(*services.UserService)
					id := params.Args["id"].(int)
					userInput := params.Args["user"].(map[string]interface{})
					email := userInput["email"].(string)
					username := userInput["username"].(string)
					password := userInput["password"].(string)
					tel, _ := userInput["tel"].(string)
					role, _ := userInput["role"].(string)

					user, err := userService.GetUserByID(id)
					if err != nil {
						return nil, err
					}

					user.Email = email
					user.Username = username
					user.Password = password
					user.Tel = tel
					user.Role = role

					err = userService.UpdateUser(user)
					if err != nil {
						return nil, err
					}

					return user, nil
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

func NewSchema(userService *services.UserService) (*graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
