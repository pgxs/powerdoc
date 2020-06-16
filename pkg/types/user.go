package types

import (
	"time"

	"pgxs.io/chassis"
)

//UserDO user database object/struct
type UserDO struct {
	UID            string
	Username       string
	Name           string
	Avatar         string
	Email          string
	Password       string
	LastLoggedInAt time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

//UserDTO  data transfer object for user service
type UserDTO struct {
}

//UserPageDTO user page
type UserPageDTO struct {
	chassis.Page
}
