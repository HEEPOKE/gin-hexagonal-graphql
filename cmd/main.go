package main

import (
	"log"

	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/http"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
	"github.com/gin-gonic/gin"
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

	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)

	server := http.NewServer(*config.Cfg)

	server.ConfigureRoutes()
	server.Start()
}
