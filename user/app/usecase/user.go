package usecase

import (
	"context"

	"github.com/isy/grpc-sandbox/user/app/domain/model"
	"github.com/isy/grpc-sandbox/user/app/domain/repository"
)

type (
	User interface {
		List(context.Context) ([]*model.User, error)
	}

	user struct {
		repo repository.User
	}
)

func NewUser(repo repository.User) User {
	return &user{
		repo: repo,
	}
}

func (u user) List(_ context.Context) ([]*model.User, error) {
	return u.repo.FindUsers()
}
