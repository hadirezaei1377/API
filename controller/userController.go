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


	// when we are binding data we must validate

if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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

	// jwt and middlewares
	func LoginUser(c echo.Context) error {
	loginModel := new(userViewModel.LoginUserViewModel)

	if err := c.Bind(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "Model not Valid")
	}
}


*/
