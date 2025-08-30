package repository

import (
	"context"

	"github.com/google/uuid"
)

type (
	UserRepoImpl interface {
		RegisterUser(ctx context.Context, chatID int64, usermane string, isAdmin bool) error
		// GetUserByChatID(ctx context.Context, chatID int64) (*domain.User, error)
		// GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
		// IsAdmin(ctx context.Context, chatID int64) bool
		// DeleteUser(ctx context.Context, username string) error
	}
	DebtRepoImpl interface {
		CreateDebt(ctx context.Context, userID uuid.UUID, amount float64, desc string) error
		// GetUserDebts(ctx context.Context, username string) ([]domain.Debt, error)
		// GetMyDebts(ctx context.Context, chatID int64) ([]domain.Debt, error)
		// MakePayment(ctx context.Context, id int64, amount float64) error
		// GetMyDebtSummary(ctx context.Context, chatID int64) (*domain.Debt, error)
	}
	// AdminRepoImpl interface {
	// 	GetAllUsers(ctx context.Context) ([]domain.User, error)
	// 	GetAllDebts(ctx context.Context) ([]domain.Debt, error)
	// 	GetUserDebtsAdmin(ctx context.Context, user entity.User) ([]domain.Debt, error)
	// 	GetDebtSummaryAdmin(ctx context.Context, user entity.User) ([]domain.Debt, error)
	// 	UpdateDebt(ctx context.Context, debt entity.Debt, amount float64) error
	// 	DeleteDebt(ctx context.Context, debt entity.Debt) error
	// }
)
