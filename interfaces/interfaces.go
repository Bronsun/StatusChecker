package interfaces

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Login    string
	Email    string
	Password string
}

type UserLink struct {
	gorm.Model
	Link   string
	Status string
	Time   time.Time
	UserId uint
}

type ResponseLink struct {
	ID     uint
	Link   string
	Status string
	Time   time.Time
}

type ResponseUser struct {
	ID    uint
	Login string
	Email string
	Links []ResponseLink
}

type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
