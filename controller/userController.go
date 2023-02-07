package controller

/*

func GetUserList(c echo.Context) error {

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, userList)
}
func CreateNewUser(c echo.Context) error {
	newUser := new(userViewModel.CreateNewUserViewModel)

	// ** bind info from client to that struct
	if err := c.Bind(newUser); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	userService := service.NewUserService()
	newUserId, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newUserId,
	}

	return c.JSON(http.StatusOK, userResData)
}


*/
