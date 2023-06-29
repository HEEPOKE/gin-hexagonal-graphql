package schemas

import "github.com/graphql-go/graphql"

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
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
			"deleted_at": &graphql.Field{
				Type: graphql.String,
			},
			"owner_id": &graphql.Field{
				Type: graphql.Int,
			},
			"owner": &graphql.Field{
				Type: UserType,
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
