package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_health "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/isy/grpc-sandbox/user/app/infra/dao"
	ui_grpc "github.com/isy/grpc-sandbox/user/app/presentation/grpc"
	"github.com/isy/grpc-sandbox/user/app/usecase"
	"github.com/isy/grpc-sandbox/user/lib/logger"
	pb "github.com/isy/grpc-sandbox/user/pb/user"
)

func main() {
	// DI
	userRepo := dao.NewUser()
	userUseCase := usecase.NewUser(userRepo)
	grcpUserUI := ui_grpc.NewUser(userUseCase)

	grpcHealthUI := ui_grpc.NewHealth()

	// lib
	if err := logger.NewLogger(); err != nil {
		fmt.Printf("Initialization error of zap logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":8080")
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_zap.UnaryServerInterceptor(logger.GetLogger()),
		),
	)

	pb.RegisterUserServiceServer(s, grcpUserUI)
	grpc_health.RegisterHealthServer(s, grpcHealthUI)

	reflection.Register(s)

	if err != nil {
		fmt.Printf("network I/O error: %v", err)
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
