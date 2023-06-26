package main

import (
	"log"

	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/repositories"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/http"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// @title			Swagger Example API
// @title			Go Gin Graphql Hexagonal API
// @version		1.0
// @description	This is a Go Gin Graphql Hexagonal API server.
// @host			localhost:6476
// @BasePath		/apis
func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)

	userRepository := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepository)

	schema, err := graphql.NewSchema(userService)
	if err != nil {
		panic(err)
	}

	router.POST("/graphql", func(c *gin.Context) {
		handler.GraphQLHandler(schema, c.Writer, c.Request)
	})

	server := http.NewServer(*config.Cfg)

	server.ConfigureRoutes()
	server.Start()
}
