package routing

import (
	"API/controller"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {

	g := e.Group("users")

	g.GET("/getList", controller.GetUserList)

	return nil
}
