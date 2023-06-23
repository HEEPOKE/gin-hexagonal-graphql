package http

import (
	"fmt"
	"time"

	_ "github.com/HEEPOKE/gin-hexagonal-graphql/internal/app/docs"
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/core/utils"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	config config.Config
}

func NewServer(config config.Config) *Server {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(helmet.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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
