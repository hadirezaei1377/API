package controller

import (
	"API/Utility"
	"API/ViewModel/common/httpResponse"
	userViewModel "API/ViewModel/user"
	"API/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetUserList(c echo.Context) error
	CreateNewUser(c echo.Context) error
	EditUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	EditUserRole(c echo.Context) error
	EditUserPassword(c echo.Context) error
}

type userController struct {
}

func NewUserController() UserController {
	return userController{}
}

func (uc userController) GetUserList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userList))
}
func (uc userController) CreateNewUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	operatorUserId, err := apiContext.GetUserId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	userService := service.NewUserService()
	isValid := userService.IsUserValidForAccess(operatorUserId, "CreateUser")
	if !isValid {
		return c.JSON(http.StatusForbidden, "")
	}

	newUser := new(userViewModel.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	newUser.CreatorUserId = operatorUserId

	newUserId, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newUserId,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

func (uc userController) EditUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUser(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) DeleteUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.DeleteUser(targetUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) EditUserRole(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserRoleViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUserRole(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) EditUserPassword(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserPasswordViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUserPassword(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}
