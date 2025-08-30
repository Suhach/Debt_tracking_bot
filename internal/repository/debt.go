package repository

import (
	"context"

	"github.com/Suhach/Debt_tracking_bot/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type (
	DebtRepoStr struct {
		db *pgxpool.Pool
	}
)

func NewDebtRepo(db *pgxpool.Pool) *DebtRepoStr {
	return &DebtRepoStr{db: db}
}

func (r *DebtRepoStr) CreateDebt(ctx context.Context, userID uuid.UUID, amount float64, desc string) error {
	query := `
		INSERT INTO debts (user_id, amount, description)
		VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, userID, amount, desc)
	if err != nil {
		logger.Log.Error("Failed to create debt", zap.Error(err))
		return err
	}
	return nil
}
