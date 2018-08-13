package graphql

import "testing"
import (
	"github.com/graphql-go/graphql"
	gql "zeego/graphql"
	"github.com/stretchr/testify/assert"
)

func TestUserQuery(t *testing.T) {
	query := `{
		user(userId: "alice") {
			id
			firstName
		}
	}`

	result := graphql.Do(graphql.Params{
		Schema:        gql.ZevereSchema,
		RequestString: query,
	})

	assert.NotNil(t, result)
	assert.Nil(t, result.Errors)
	assert.NotNil(t, result.Data)

	data := result.Data.(map[string]interface{})
	assert.NotNil(t, data["user"])

	user := data["user"].(map[string]interface{})
	assert.Equal(t, "alice", user["id"])
	assert.Equal(t, "Alice", user["firstName"])
	assert.Equal(t, 2, len(user))
}
