package resolver

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/graphql-go/graphql"
)

type RootResolver struct {
	UserResolver *UserResolver
}

func NewRootResolver(userService *services.UserService) *RootResolver {
	return &RootResolver{
		UserResolver: &UserResolver{
			UserService: userService,
		},
	}
}

func (r *RootResolver) SetupResolvers(schema graphql.Schema) {
	schemaConfig := schema.Config

	if schemaConfig.Query != nil {
		queryFields := schemaConfig.Query.(*graphql.Object).Fields
		queryFields["users"].Resolve = r.UserResolver.GetAllUsers
		queryFields["user"].Resolve = r.UserResolver.GetUserByID
	}

	if schemaConfig.Mutation != nil {
		mutationFields := schemaConfig.Mutation.(*graphql.Object).Fields
		mutationFields["createUser"].Resolve = r.UserResolver.CreateUser
		mutationFields["updateUser"].Resolve = r.UserResolver.UpdateUser
		mutationFields["deleteUser"].Resolve = r.UserResolver.DeleteUser
	}
}
