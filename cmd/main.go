package main

import (
	"log"

	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/http"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}

	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect to the database: %s", err)
	}

	server := http.NewServer(*cfg)
	server.ConfigureGraphQLRoutes()
	server.Start()
}
