package main

import (
	"log"

	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/http"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/database"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServer(*config.Cfg)
	server.ConfigureGraphQLRoutes()
	server.Start()
}
