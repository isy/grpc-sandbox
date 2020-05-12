package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	// "github.com/grpc-ecosystem/go-grpc-middleware"

	pb "github.com/isy/grpc-sandbox/user/pb/user"
)

type userService struct{}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	s := grpc.NewServer(
	// grpc.StreamInterceptor()
	)

	us := &userService{}

	pb.RegisterUserServiceServer(s, us)

	reflection.Register(s)

	if err != nil {
		fmt.Println("network I/O error: %v", err)
		os.Exit(1)
	}

	if err := s.Serve(lis); err != nil {
		fmt.Printf("serve error: %v", err)
		os.Exit(1)
	}
	fmt.Println("Run server")
}

func (s *userService) ListUsers(ctx context.Context, in *emptypb.Empty) (*pb.ListUsersResponse, error) {
	fmt.Println("test")

	return &pb.ListUsersResponse{
		Users: []*pb.User{
			{Id: 1},
		},
	}, nil
}
