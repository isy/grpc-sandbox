package service

import (
	"fmt"

	"github.com/isy/grpc-sandbox/gateway/app/service/user"
)

type Client struct {
	UserClient *user.UserClient
}

func BindGrpcClient() *Client {
	uc, err := user.NewClient()
	if err != nil {
		fmt.Printf("Connection errors in user services: %v", err)
	}

	return &Client{
		UserClient: uc,
	}
}

func (c Client) Close() {
	c.UserClient.Conn.Close()
}
