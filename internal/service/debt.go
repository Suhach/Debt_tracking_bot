package service

import (
	"context"

	"github.com/Suhach/Debt_tracking_bot/domain"
	"github.com/Suhach/Debt_tracking_bot/internal/repository"
)

type (
	DebtServiceStr struct {
		repo repository.DebtRepoImpl
	}
)

func NewDebtService(repo repository.DebtRepoImpl) *DebtServiceStr {
	return &DebtServiceStr{repo: repo}
}

func (s *DebtServiceStr) CreateDebt(ctx context.Context, debt domain.Debt) error {
	return s.repo.CreateDebt(ctx, debt.UserID, debt.Amount, debt.Description)
}
