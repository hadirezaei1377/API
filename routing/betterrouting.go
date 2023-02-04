package routing

/*

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {

	// using groups instead of each parameters like users
	g := e.Group("users")

	myfunction := func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/   with param")
	}

	g.GET("/get/:id", myfunction).Name = "getUserById"
	g.GET("/get/:id/avatars", getUserAvatar)
	g.GET("/get/:id/avatars/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/[with param]/avatars/*")
	})
	g.GET("/get/15", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/15")
	}).Name = "getUserNo15" // the name that we want

	for i, route := range e.Routes() {
		fmt.Println(i, route.Method, route.Path, route.Name)
	}

	fmt.Println(e.URI(getUserAvatar, 17))
	fmt.Println(e.Reverse("getUserNo15", nil))
	fmt.Println(e.Reverse("getUserById", 13))

	return nil
}
func getUserAvatar(c echo.Context) error {
	return c.String(http.StatusOK, "/users/get/[with param]/avatars")
}

*/
