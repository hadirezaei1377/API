package service

/*

// service for improrve the Project architecture

import "API/model/user"

// an user can user interface for using interface but not the main userservice

type UserService interface {
	GetUserList() ([]user.User, error)
}

// private struct for main userservice
// controller and other packages dont have an access to it
type userService struct {
}

// return new user service
func NewUserService() UserService {
	return userService{}
}

// function for organize the list of users instead controller
func (userService) GetUserList() ([]user.User, error) {
	userList := []user.User{
		{
			FirstName:   "Hadi",
			LastName:    "Rezaei",
			Age:         24,
			PhoneNumber: "09131111111",
		},
		{
			FirstName:   "Khashayar",
			LastName:    "Rezaei",
			Age:         35,
			PhoneNumber: "09911111111",
		},
		{
			FirstName:   "Gisoo",
			LastName:    "Sotoodeh",
			Age:         21,
			PhoneNumber: "09351111111",
		},
	}

	return userList, nil
}


*/

// bind user

/*


type UserService interface {
	GetUserList() ([]user.User, error)
	CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error)
	GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (userService) GetUserList() ([]user.User, error) {

	userRepository := repository.NewUserRepository()
	userList, err := userRepository.GetUserList()

	return userList, err
}

// related to authentication
func (userService) GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error) {

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserByUserNameAndPassword(loginViewModel.UserName, loginViewModel.Password)

	return user, err
}


func (userService) CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error) {

	userEntity := user.User{
		FirstName:    userInput.FirstName,
		LastName:     userInput.LastName,
		Email:        userInput.Email,
		UserName:     userInput.UserName,
		Password:     userInput.Password,
		RegisterDate: time.Now(),
	}

	userRepository := repository.NewUserRepository()
	userId, err := userRepository.InsertUser(userEntity)

	return userId, err
}

*/
