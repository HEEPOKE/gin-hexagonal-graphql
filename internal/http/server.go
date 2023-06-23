package http

import (
	"fmt"

	_ "github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/docs"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/core/utils"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		router: gin.Default(),
		config: config,
	}
}

func (s *Server) ConfigureRoutes() {
	api := s.router.Group("/apis")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api.GET("/", utils.HandleFirst)
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.config.PORT)
	s.router.Run(addr)
}
