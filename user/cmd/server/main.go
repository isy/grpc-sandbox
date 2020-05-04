package main

import (
	"google.golang.org/grpc"
	pb "github.com/isy/grpc-sandbox/user/interface/grpc"
)

type Service struct {}

func main() {
	server := grpc.NewServer()

	s := &Service{}

	pb.RegisterUserServiceServer(server, s)
}

func (*Service) ListUsers(ctx context.Context) (*pb.ListUsersResponse, error) {
	fmt.Println("test")

	return &pb.ListUsersResponse {

	},
}