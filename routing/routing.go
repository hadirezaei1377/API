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

	userController := controller.NewUserController()
	accountController := controller.NewAccountController()
	newsController := controller.NewNewsController()

	e.POST("/login", accountController.LoginUser)
	e.POST("/UploadAvatar", userController.UploadAvatar)

	g := e.Group("users")

	g.GET("/getList", userController.GetUserList)

	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     &security.JwtClaims{},
	}
	g.POST("/CreateNewUser", userController.CreateNewUser, PermissionChecker("CreateUser"), middleware.JWTWithConfig(jwtConfig))

	g.PUT("/EditUser/:id", userController.EditUser, PermissionChecker("EditUser"), middleware.JWTWithConfig(jwtConfig))
	g.DELETE("/DeleteUser/:id", userController.DeleteUser, PermissionChecker("DeleteUser"), middleware.JWTWithConfig(jwtConfig))

	g.PUT("/EditUserRole/:id", userController.EditUserRole, middleware.JWTWithConfig(jwtConfig))
	g.PUT("/EditUserPassword/:id", userController.EditUserPassword, middleware.JWTWithConfig(jwtConfig))

	newsGroup := e.Group("news")

	newsGroup.GET("/getList", newsController.GetNewsList)

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
