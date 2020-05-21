package user

import (
	"fmt"
	"net/http"

	"github.com/isy/grpc-sandbox/gateway/app/aggregater"
	"github.com/labstack/echo/v4"
)

type (
	User interface {
		Users(c echo.Context) error
	}

	user struct {
		ag aggregater.User
	}
)

func NewUser(ag aggregater.User) User {
	return &user{
		ag: ag,
	}
}

func (u user) Users(c echo.Context) error {
	users, err := u.ag.Users(c.Request().Context())
	if err != nil {
		fmt.Printf("request fail: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}
