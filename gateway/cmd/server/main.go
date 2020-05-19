package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/isy/grpc-sandbox/gateway/lib/logger"
	pb_user "github.com/isy/grpc-sandbox/gateway/pb/user"
)

func main() {
	// lib
	if err := logger.NewLogger(); err != nil {
		fmt.Printf("Initialization error of zap logger: %v", err)
	}

	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
		log.Fatal("client connection err:", err)
	}
	defer conn.Close()

	client := pb_user.NewUserServiceClient(conn)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Running!!")
	})

	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", func(c echo.Context) error {
				message := &emptypb.Empty{}
				res, err := client.ListUsers(context.Background(), message)
				if err != nil {
					log.Fatal("request fail:", err)
					return c.JSON(http.StatusInternalServerError, err)
				}
				return c.JSON(http.StatusOK, res)
			})
		}
	}

	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Print("shutdown server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
