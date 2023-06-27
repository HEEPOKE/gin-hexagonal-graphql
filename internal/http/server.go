package http

import (
	"fmt"
	"time"

	ConfigGraphql "github.com/HEEPOKE/gin-hexagonal-graphql/internal/server/graphql"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
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
		router: router,
		config: config,
	}
}

func (s *Server) ConfigureGraphQLRoutes() {
	h := handler.New(&handler.Config{
		Schema:   ConfigGraphql.GetSchema(),
		Pretty:   true,
		GraphiQL: true,
	})

	api := s.router.Group("/api")
	api.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.config.PORT)
	s.router.Run(addr)
}
