package resolver

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/graphql-go/graphql"
)

type UserResolver struct {
	UserService *services.UserService
}

func NewUserResolver(userService *services.UserService) *UserResolver {
	return &UserResolver{
		UserService: userService,
	}
}

func (r *UserResolver) ResolveGetAllUsers(params graphql.ResolveParams) (interface{}, error) {
	users, err := r.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserResolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) ResolveCreateUser(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)
	email := params.Args["email"].(string)

	user := &models.User{
		Username: name,
		Email:    email,
	}

	err := r.UserService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) ResolveUpdateUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	name := params.Args["name"].(string)
	email := params.Args["email"].(string)

	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Username = name
	user.Email = email

	err = r.UserService.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) ResolveDeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	err = r.UserService.DeleteUser(user)
	if err != nil {
		return nil, err
	}

	return true, nil
}
