package service

import (
	"context"

	"github.com/bildanjhry/log-be/internal/model"
	"github.com/bildanjhry/log-be/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.UserRes, error) {
	return s.repo.FindAll(ctx)
}
