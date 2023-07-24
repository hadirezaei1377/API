package routing

import (
	"API/Utility"
	"API/ViewModel/common/security"
	"API/controller"
	"API/service"
	"net/http"

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
	g.POST("/CreateNewUser", controller.CreateNewUser, PermissionChecker("CreateNewUser"), middleware.JWTWithConfig(jwtConfig))

	return nil
}
func PermissionChecker(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := c.(*Utility.ApiContext)

			operatorUserId, err := apiContext.GetUserId()
			if err != nil {
				return &echo.HTTPError{
					Code:     401,
					Message:  http.StatusUnauthorized,
					Internal: err,
				}
			}

			userService := service.NewUserService()
			isValid := userService.IsUserValidForAccess(operatorUserId, permission)
			if !isValid {
				return &echo.HTTPError{
					Code:    403,
					Message: http.StatusForbidden,
				}
			}

			return next(c)
		}
	}
}
