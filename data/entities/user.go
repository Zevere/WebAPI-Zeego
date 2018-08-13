package entities

import (
	"time"
)

type User struct {
	Id             string
	Name           string
	PassphraseHash string
	FirstName      string
	LastName       *string
	JoinedAt       time.Time
	ModifiedAt     *time.Time
	DeletedAt      *time.Time
}
