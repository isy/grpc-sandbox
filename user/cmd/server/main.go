package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/isy/grpc-sandbox/user/app/infra/dao"
	ui_grpc "github.com/isy/grpc-sandbox/user/app/presentation/grpc"
	"github.com/isy/grpc-sandbox/user/app/usecase"
	pb "github.com/isy/grpc-sandbox/user/pb/user"
)

func main() {
	// DI
	userRepo := dao.NewUser()
	userUseCase := usecase.NewUser(userRepo)
	grcpUserUI := ui_grpc.NewUser(userUseCase)

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

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("serve error: %v", err)
			os.Exit(1)
		}
	}()
	defer s.GracefulStop()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
