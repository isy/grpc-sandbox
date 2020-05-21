package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/isy/grpc-sandbox/gateway/app/handler/rest/user"
)

type Handler struct {
	User user.User
}

func Bind(e *echo.Echo, h *Handler) {

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Running!!")
	})

	api := e.Group("/v1")
	users := api.Group("/users")
	{
		users.GET("/", h.User.Users)
	}
}
