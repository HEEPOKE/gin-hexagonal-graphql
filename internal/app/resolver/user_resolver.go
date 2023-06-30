package resolver

import (
	"time"

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
	return users, err
}

func (r *UserResolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	user, err := r.UserService.GetUserByID(id)
	return user, err
}

func (r *UserResolver) ResolveCreateUser(params graphql.ResolveParams) (interface{}, error) {
	username := params.Args["username"].(string)
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)
	tel := params.Args["tel"].(string)
	role := params.Args["role"].(string)

	user := &models.User{
		Username:  username,
		Email:     email,
		Password:  password,
		Tel:       tel,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.UserService.CreateUser(user)
	return user, err
}

func (r *UserResolver) ResolveUpdateUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	username := params.Args["username"].(string)
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)
	tel := params.Args["tel"].(string)
	role := params.Args["role"].(string)

	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email
	user.Password = password
	user.Tel = tel
	user.Role = role
	user.UpdatedAt = time.Now()

	_, err = r.UserService.UpdateUser(id, user)
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
	return true, err
}
