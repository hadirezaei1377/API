package routing

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {

	// What address did I get, what should I do?
	e.GET("/users/get/:id", func(c echo.Context) error {
		// send this response to client
		return c.String(http.StatusOK, "/users/get/   with param")
		// with param shows to client the using part
	})

	e.GET("/users/get/:id/avatars", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/[with param]/avatars")
	})

	// for any params that client sends :
	e.GET("/users/get/:id/avatars/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/[with param]/avatars/*")
	})

	e.GET("/users/get/15", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/15")
	})

	// print routs
	for i, route := range e.Routes() {
		fmt.Println(i, route.Method, route.Path)
	}

	return nil

}
