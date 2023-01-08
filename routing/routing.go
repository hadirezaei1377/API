package routing

import (
	"API/controller"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {

	g := e.Group("users")

	g.GET("/get/:id/avatars", controller.GetUserAvatar)

	g.GET("/get", controller.GetUser)
	g.POST("/create", controller.CreateUser)

	return nil
}
