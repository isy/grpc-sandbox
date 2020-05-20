package user

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"

	"github.com/isy/grpc-sandbox/gateway/lib/logger"
	pb_user "github.com/isy/grpc-sandbox/gateway/pb/user"
)

type UserClient struct {
	Conn   *grpc.ClientConn
	Client pb_user.UserServiceClient
}

func NewClient() (*UserClient, error) {
	conn, err := grpc.Dial(
		"user-service:8080",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				grpc_zap.UnaryClientInterceptor(logger.GetLogger()),
			),
		),
	)
	if err != nil {
		return nil, err
	}

	client := pb_user.NewUserServiceClient(conn)

	return &UserClient{
		Conn:   conn,
		Client: client,
	}, nil
}
