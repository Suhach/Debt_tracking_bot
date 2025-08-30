package service

import (
	"context"

	"github.com/Suhach/Debt_tracking_bot/domain"
)

type (
	DebtServiceImpl interface {
		CreateDebt(ctx context.Context, debt domain.Debt) error
		// GetUserDebts(ctx context.Context, debt domain.Debt)
	}

	UserServiceImpl interface {
		CreateUser(ctx context.Context, user domain.User) error
	}

	// AdminServiceImpl interface {
	// 	GetAllUsers(ctx context.Context) ([]dto.GetUserResponse, error)
	// }
)
