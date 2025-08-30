package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	ChatID    int64
	Username  string
	CreatedAt time.Time
	IsAdmin   bool
}

func (u *User) Isdmin() bool {
	if u.IsAdmin == true {
		return true
	}
	return false
}
