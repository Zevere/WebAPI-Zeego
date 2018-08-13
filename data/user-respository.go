package data

import (
	"zeego/data/entities"
	"io"
)

type UserRepository interface {
	io.Closer
	Insert(u *entities.User) Error
	GetByName(name string) (*entities.User, Error)
}
