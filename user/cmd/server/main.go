package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/isy/grpc-sandbox/user/app/interface/grpc"
)

type Service struct{}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	server := grpc.NewServer()

	s := &Service{}

	pb.RegisterUserServiceServer(server, s)

	reflection.Register(server)

	if err != nil {
		fmt.Println("network I/O error: %v", err)
		os.Exit(1)
	}

	if err := server.Serve(lis); err != nil {
		fmt.Printf("serve error: %v", err)
		os.Exit(1)
	}
	fmt.Println("Run server")
}

func (s *Service) ListUsers(ctx context.Context, in *emptypb.Empty) (*pb.ListUsersResponse, error) {
	fmt.Println("test")

	return &pb.ListUsersResponse{
		Users: []*pb.User{
			{Id: 1},
		},
	}, nil
}
