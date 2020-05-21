package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/isy/grpc-sandbox/gateway/app/aggregater/repository"
	pb_user "github.com/isy/grpc-sandbox/gateway/pb/user"
)

type user struct {
	uc pb_user.UserServiceClient
}

func NewUser(uc pb_user.UserServiceClient) repository.User {
	return &user{
		uc: uc,
	}
}

func (u user) UserList(ctx context.Context) (*pb_user.ListUsersResponse, error) {
	message := &emptypb.Empty{}
	res, err := u.uc.ListUsers(ctx, message)
	if err != nil {
		return nil, err
	}

	return res, nil
}
