package adapters

import (
	"context"

	"github.com/Suhach/Debt_tracking_bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	DebtAdapterStr struct {
		srv service.DebtServiceImpl
	}
)

func NewDebtAdapter(srv service.DebtServiceImpl) *DebtAdapterStr {
	return &DebtAdapterStr{srv: srv}
}

func (a *DebtAdapterStr) CreateDebt(ctx context.Context, bot *tgbotapi.Message) {

}
