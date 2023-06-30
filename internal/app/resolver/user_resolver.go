package resolver

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models/response"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/constants"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
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
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(users, constants.SUCCESS)
	return successResponse, nil
}

func (r *UserResolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			notFoundResponse := response.NewResponse(nil, constants.NOT_FOUND)
			return notFoundResponse, nil
		}

		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(user, constants.SUCCESS)
	return successResponse, nil
}

func (r *UserResolver) ResolveCreateUser(params graphql.ResolveParams) (interface{}, error) {
	userInput := params.Args["user"].(map[string]interface{})

	email := userInput["email"].(string)
	username := userInput["username"].(string)
	password := userInput["password"].(string)
	tel := userInput["tel"].(string)
	role, _ := userInput["role"].(string)

	user := &models.User{
		Email:    email,
		Username: username,
		Password: password,
		Tel:      tel,
		Role:     role,
	}

	err := r.UserService.CreateUser(user)
	if err != nil {
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(user, constants.SUCCESS)
	return successResponse.Payload, nil
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
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	user.ID = id
	user.Username = username
	user.Email = email
	user.Password = password
	user.Tel = tel
	user.Role = role

	_, err = r.UserService.UpdateUser(user)
	if err != nil {
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(user, constants.SUCCESS)
	return successResponse, nil
}

func (r *UserResolver) ResolveDeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			notFoundResponse := response.NewResponse(nil, constants.NOT_FOUND)
			return notFoundResponse, nil
		}

		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	err = r.UserService.DeleteUser(user)
	if err != nil {
		errorResponse := response.NewResponse(nil, constants.FAILED)
		errorResponse.Status.Error = err.Error()
		return errorResponse, nil
	}

	successResponse := response.NewResponse(nil, constants.SUCCESS)
	return successResponse, nil
}
