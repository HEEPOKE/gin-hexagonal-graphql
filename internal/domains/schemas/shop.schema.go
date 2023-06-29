package schemas

import (
	"github.com/HEEPOKE/gin-hexagonal-graphql/internal/domains/models"
	"github.com/graphql-go/graphql"
)

var ShopType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Shop",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"telShop": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					shop := p.Source.(*models.Shop)
					return shop.CreatedAt, nil
				},
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					shop := p.Source.(*models.Shop)
					return shop.UpdatedAt, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					shop := p.Source.(*models.Shop)
					return shop.DeletedAt.Time, nil
				},
			},
			"owner_id": &graphql.Field{
				Type: graphql.Int,
			},
			"owner": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					shop := p.Source.(*models.Shop)
					return shop.Owner, nil
				},
			},
		},
	},
)

var ShopInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ShopInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"location": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"telShop": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"owner_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
	},
)
