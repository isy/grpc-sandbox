package aggregater

import (
	"context"

	"github.com/isy/grpc-sandbox/gateway/app/aggregater/dto"
	"github.com/isy/grpc-sandbox/gateway/app/aggregater/repository"
)

type (
	User interface {
		Users(ctx context.Context) ([]*dto.User, error)
	}

	user struct {
		ur repository.User
	}
)

func NewUser(ur repository.User) User {
	return &user{
		ur: ur,
	}
}

func (u user) Users(ctx context.Context) ([]*dto.User, error) {
	users, err := u.ur.UserList(ctx)
	if err != nil {
		return nil, err
	}

	res := []*dto.User{}
	for _, u := range users.GetUsers() {
		res = append(res, &dto.User{
			ID:        u.Id,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Gender:    int32(u.Gender),
			CreatedAt: u.CreatedAt,
		})
	}

	return res, nil
}
