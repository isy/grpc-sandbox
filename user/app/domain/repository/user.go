package repository

import (
	"github.com/isy/grpc-sandbox/user/app/domain/model"
)

type User interface {
	FindUsers() ([]*model.User, error)
}
