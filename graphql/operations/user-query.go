package operations

import (
	"zeego/graphql/types"
	"zeego/data"
)

func GetUserById(repo *data.UserRepository, userId string) (*types.User) {
	entity := (*repo).Get(userId)
	return &types.User{
		Id:        entity.DisplayID,
		FirstName: entity.FirstName,
	}
}
