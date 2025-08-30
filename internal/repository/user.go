package repository

import (
	"context"

	"github.com/Suhach/Debt_tracking_bot/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type (
	UserSTR struct {
		db *pgxpool.Pool
	}
)

func NewUser(db *pgxpool.Pool) *UserSTR {
	return &UserSTR{db: db}
}

func (r *UserSTR) RegisterUser(ctx context.Context, chatID int64, usermane string, isAdmin bool) error {
	query := `
		INSERT INTO users (chat_id, username, is_admin)
		VALUES ($1, $2, $3)
		ON CONFLICT (chat_id) DO NOTHING`
	_, err := r.db.Exec(ctx, query, chatID, usermane, isAdmin)
	if err != nil {
		logger.Log.Error("Failed to register user", zap.Error(err))
		return err
	}
	return nil
}
