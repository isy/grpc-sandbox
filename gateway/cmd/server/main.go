package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/isy/grpc-sandbox/gateway/app/aggregater"
	"github.com/isy/grpc-sandbox/gateway/app/handler/rest"
	user_ui "github.com/isy/grpc-sandbox/gateway/app/handler/rest/user"
	"github.com/isy/grpc-sandbox/gateway/app/infra/grpc"
	"github.com/isy/grpc-sandbox/gateway/app/infra/grpc/user"
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

	client := grpc.BindGrpcClient()
	defer client.Close()

	// DI
	userRepo := user.NewUser(client.UserClient.Client)
	userAg := aggregater.NewUser(userRepo)
	userUI := user_ui.NewUser(userAg)

	handler := &rest.Handler{
		User: userUI,
	}
	rest.Bind(e, handler)

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
