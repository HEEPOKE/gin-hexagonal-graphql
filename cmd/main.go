package main

import (
	"log"

	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Run(":" + config.Cfg.PORT)
}
