package dao

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/isy/grpc-sandbox/user/app/domain/model"
	"github.com/isy/grpc-sandbox/user/app/domain/repository"
)

type user struct {
}

func NewUser() repository.User {
	return &user{}
}

func (u user) FindUsers() ([]*model.User, error) {
	return []*model.User{
		{ID: 1, FirstName: "クマ", LastName: "太郎", Gender: 1, CreatedAt: &timestamp.Timestamp{}},
	}, nil
}
