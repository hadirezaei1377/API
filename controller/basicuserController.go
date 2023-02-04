package controller

/*

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// using functions in routing here
func GetUserAvatar(c echo.Context) error {
	// c is context
	idString := c.Param("id")
	return c.String(http.StatusOK, "/users/get/[with param ("+idString+")]/avatars")
}

// requests recieve in this format
type UserSearchInput struct {
	// Name in query that we recieve is named name ,for example
	Name   string `query:"name" json:"firstName"`
	Family string `query:"family" json:"lastname"`
	Age    int    `query:"age"`
	Phone  string `query:"phone"`
}

// function for searching on users
// any received requests goes to get and then GetUser as query params
func GetUser(c echo.Context) error {

	// put request in this
	userInput := new(UserSearchInput)
	// bind those in userinput
	err := c.Bind(userInput)
	if err != nil {
		return err
	}
	fmt.Println(userInput)
	return c.String(http.StatusOK, "/users/get")
}

func CreateUser(c echo.Context) error {

	userInput := new(UserSearchInput)
	err := c.Bind(userInput)
	if err != nil {
		return err
	}
	fmt.Println(userInput)
	return c.String(http.StatusOK, "/users/CreateUser")
}

*/
