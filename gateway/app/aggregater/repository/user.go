package repository

import (
	"context"

	pb_user "github.com/isy/grpc-sandbox/gateway/pb/user"
)

type User interface {
	UserList(ctx context.Context) (*pb_user.ListUsersResponse, error) // FIXME: Respose type
}
