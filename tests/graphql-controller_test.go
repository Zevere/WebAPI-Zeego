package tests

import (
	"testing"
	_ "testy/routers"

	"zeego/controllers"
	"github.com/graphql-go/graphql"
	"github.com/astaxie/beego/context"
)

var testSchema graphql.Schema

func init() {
	config := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"foo": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "bar", nil
				},
			},
		},
	}
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(config),
	})
	testSchema = schema
}

func TestPost(t *testing.T) {
	c := controllers.NewGraphQLController(&testSchema)
	c.Ctx = &context.Context{Input: &context.BeegoInput{}}

	c.Ctx.Input.RequestBody = []byte(`{ foo }`)
	//
	//c.Post()
	//
	//assert.NotNil(t, c.Data["json"])
	//assert.IsType(t, graphql.Result{}, c.Data["json"])
	//assert.Equal(t, `{"data":{"foo":}}`, c.Data["json"])
}
