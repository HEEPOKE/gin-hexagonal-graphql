package http

import (
	"fmt"
	"log"
	"net/http"

	ConfigGraphql "github.com/HEEPOKE/gin-hexagonal-graphql/internal/server/graphql"
	"github.com/HEEPOKE/gin-hexagonal-graphql/pkg/config"
	"github.com/graphql-go/handler"
)

type Server struct {
	router *http.ServeMux
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		router: http.NewServeMux(),
		config: config,
	}
}

func (s *Server) ConfigureGraphQLRoutes() {
	var playground bool
	if s.config.LOCAL {
		playground = true
	} else {
		playground = false
	}

	h := handler.New(&handler.Config{
		Schema:     ConfigGraphql.GetSchema(),
		Pretty:     true,
		GraphiQL:   false,
		Playground: playground,
	})

	s.router.HandleFunc("/apis/graphql", func(w http.ResponseWriter, r *http.Request) {
		if !s.config.LOCAL && r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.config.PORT)
	log.Printf("Server is running at http://localhost:%d/", s.config.PORT)
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		log.Fatal(err)
	}
}
