package adapters

import "github.com/Suhach/Debt_tracking_bot/internal/service"

type (
	UserSTR struct {
		srv service.UserServiceImpl
	}
)

func NewUser(srv service.UserServiceImpl) *UserSTR {
	return &UserSTR{srv: srv}
}

func (a *UserSTR) NewUser() {
	
}
