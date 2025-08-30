package service

import (
	"context"

	"github.com/Suhach/Debt_tracking_bot/domain"
	"github.com/Suhach/Debt_tracking_bot/internal/repository"
)

type (
	UserStr struct {
		repo repository.UserRepoImpl
	}
)

func NewUserService(repo repository.UserRepoImpl) *UserStr {
	return &UserStr{repo: repo}
}
func (s *UserStr) CreateUser(ctx context.Context, user domain.User) error {
	return s.repo.RegisterUser(ctx, user.ChatID, user.Username, user.IsAdmin)
}
