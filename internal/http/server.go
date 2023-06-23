package http

import (
	"fmt"

	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/core/utils"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/gin-gonic/gin"
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
	s.router.GET("/", utils.HandleFirst)

	api := s.router.Group("/apis")
	api.GET("/", utils.HandleFirst)
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.config.PORT)
	s.router.Run(addr)
}
