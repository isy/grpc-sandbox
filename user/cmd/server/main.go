package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/isy/grpc-sandbox/user/app/infra/dao"
	"github.com/isy/grpc-sandbox/user/app/usecase"
	pb "github.com/isy/grpc-sandbox/user/pb/user"
	grpc_ui "github.com/isy/grpc-sandbox/user/presentation/grpc"
)

func main() {
	// DI
	userRepo := dao.NewUser()
	userUseCase := usecase.NewUser(userRepo)
	grcpUserUI := grpc_ui.NewUser(userUseCase)

	lis, err := net.Listen("tcp", ":8080")
	s := grpc.NewServer(
	// grpc.StreamInterceptor()
	)

	pb.RegisterUserServiceServer(s, grcpUserUI)

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
