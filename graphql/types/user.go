package types

import "github.com/graphql-go/graphql"

var UserType graphql.Type

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
}

func init() {
	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: `A user`,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "User's ID",
			},
			"firstName": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "User's first name",
			},
		},
	})
}
