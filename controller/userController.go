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


	// get claims from JWT
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*security.JwtClaims)
	newUser.CreatorUserId = claim.UserId

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





	// Authentication , before create a token  : set usernames and check passwords




	// jwt and middlewares
	func LoginUser(c echo.Context) error {
	loginModel := new(userViewModel.LoginUserViewModel)

	if err := c.Bind(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "Model not Valid")

		// we have a claims that is a jwtclaim in security
		// feed its fields(the info that exsist in log in info)
		claims := &security.JwtClaims{
		loginModel.UserName,

		// feed another field
		jwt.StandardClaims{
			// take now and add be right for 24 hours and change it to unix
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}



	// create token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// encode our token
	stringToken, err := token.SignedString([]byte("secret")) // show the tokens string
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// take created token to client
	return c.JSON(http.StatusOK, echo.Map{
		"token": stringToken,
	})
}
	}
}


*/
