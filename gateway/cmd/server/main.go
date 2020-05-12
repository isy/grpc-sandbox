package main

import (
	"context"
	"log"
	"net/http"

	// "golang.org/x/xerrors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb_user "github.com/isy/grpc-sandbox/gateway/pb/user"
)

func main() {
	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	conn, err := grpc.Dial("user-service:8080", grpc.WithInsecure())

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
				// return xerrors.New("エラー")
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

	e.Logger.Fatal(e.Start(":8080"))
}
