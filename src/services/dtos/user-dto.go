package dtos

import "time"

type UserDto struct {
	Username  string
	Password  string
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
