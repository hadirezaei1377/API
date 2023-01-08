package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserList(c echo.Context) error {

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, userList)
}
