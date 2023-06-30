package graphql

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/server"
	"github.com/graphql-go/graphql"
)

var schema graphql.Schema

func init() {
	rootQuery := graphql.ObjectConfig{
		Name:   "Query",
		Fields: server.GetRootQueryFields(),
	}

	rootMutation := graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: server.GetRootMutationFields(),
	}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	var err error
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}
}

func GetSchema() *graphql.Schema {
	return &schema
}
