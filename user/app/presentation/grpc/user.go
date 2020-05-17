package grpc

import (
	"context"

	"golang.org/x/xerrors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/isy/grpc-sandbox/user/app/usecase"
	pb "github.com/isy/grpc-sandbox/user/pb/user"
)

type (
	User interface {
		ListUsers(context.Context, *emptypb.Empty) (*pb.ListUsersResponse, error)
	}
	user struct {
		uc usecase.User
	}
)

func NewUser(usecase usecase.User) User {
	return &user{
		uc: usecase,
	}
}

func (u *user) ListUsers(ctx context.Context, in *emptypb.Empty) (*pb.ListUsersResponse, error) {
	users, err := u.uc.List(ctx)
	if err != nil {
		return nil, xerrors.New("fail")
	}

	resUser := []*pb.User{}

	for _, user := range users {
		resUser = append(resUser, &pb.User{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    pb.User_Gender(user.Gender),
			CreatedAt: user.CreatedAt,
		})
	}

	return &pb.ListUsersResponse{
		Users: resUser,
	}, nil
}