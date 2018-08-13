package models

import "time"

type User struct {
	Id             string
	DisplayId      string
	FirstName      string
	PassphraseHash string
	JoinedAt       time.Time
}
