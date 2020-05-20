package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/isy/grpc-sandbox/gateway/app/service"
	"github.com/isy/grpc-sandbox/gateway/lib/logger"
)

func main() {
	// lib
	if err := logger.NewLogger(); err != nil {
		fmt.Printf("Initialization error of zap logger: %v", err)
	}

	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	client := service.BindGrpcClient()
	defer client.Close()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Running!!")
	})

	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", func(c echo.Context) error {
				message := &emptypb.Empty{}
				res, err := client.UserClient.Client.ListUsers(context.Background(), message)
				if err != nil {
					fmt.Printf("request fail: %v", err)
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
		fmt.Printf("Shutdown error: %v", err)
	}
}
