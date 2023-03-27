package routing

/*

import (
	"API/controller"

	"github.com/labstack/echo/v4"
)

// group level middlewares
func GroupLevel(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("GroupLevel")
	return next
}

// route level middlewares
// related to get and post in bottom of this file
func RouteLevel(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("RouteLevel")
	return next
}
func RouteLevel2(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("RouteLevel2")
	return next
}
func RouteLevel3(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("RouteLevel3")
	return next
}

func SetRouting(e *echo.Echo) error {

	g := e.Group("users", GroupLevel)

	g.GET("/getList", controller.GetUserList, RouteLevel, RouteLevel2, RouteLevel3)
	g.POST("/CreateNewUser", controller.CreateNewUser)

	return nil
}

*/
