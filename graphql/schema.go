package graphql

import (
	"github.com/graphql-go/graphql"
	"log"
	"zeego/graphql/types"
	"zeego/graphql/operations"
)

var ZevereSchema graphql.Schema

func init() {
	config := graphql.ObjectConfig{
		Name:   "UserQuery",
		Fields: getUserQueryFields(),
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(config),
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
		panic("Invalid schema")
	}
	ZevereSchema = schema
}

func getUserQueryFields() graphql.Fields {
	return graphql.Fields{
		"user": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["userId"].(string)
				return operations.GetUserById(nil, id), nil // ToDo
			},
		},
	}
}
