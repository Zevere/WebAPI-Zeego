package routers

import (
	"github.com/graphql-go/graphql"
	"github.com/astaxie/beego"
	"zeego/controllers"
)

func initRoutes(
	schema *graphql.Schema,
) {
	beego.Router("/api/GraphQL", controllers.NewGraphQLController(schema))
}
