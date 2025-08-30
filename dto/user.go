package dto

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateUserRequest struct {
		ChatID   int64  `json:"chat_id" validate:"required"`
		Username string `json:"username" validate:"required"`
	}
	CreateUserResponse struct {
		ID        uuid.UUID `json:"id"`
		ChatID    int64     `json:"chat_id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	}

	GetUserResponse struct {
		ID        uuid.UUID `json:"id"`
		ChatID    int64     `json:"chat_id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	}
)
