package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	ChatID    int64     `db:"chat_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	IsAdmin   bool      `db:"is_admin"`
}
