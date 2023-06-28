package routing

import (
	"API/ViewModel/common/security"
	"API/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting(e *echo.Echo) error {

	e.POST("/login", controller.LoginUser)

	g := e.Group("users")

	g.GET("/getList", controller.GetUserList)

	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     &security.JwtClaims{},
	}
	g.POST("/CreateNewUser", controller.CreateNewUser, middleware.JWTWithConfig(jwtConfig))

	return nil
}
