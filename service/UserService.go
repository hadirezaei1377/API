package service

import (
	userViewModel "API/ViewModel/user"
	"API/model/user"
	"API/repository"
	"time"

	"golang.org/x/exp/slices"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error)
	GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error)
	IsUserValidForAccess(userId, roleName string) bool
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

func (userService) GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error) {

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserByUserNameAndPassword(loginViewModel.UserName, loginViewModel.Password)

	return user, err
}
func (userService) IsUserValidForAccess(userId, roleName string) bool {

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserById(userId)

	if err != nil {
		return false
	}

	if user.Roles == nil {
		return false
	}
	roleIndex := slices.IndexFunc(user.Roles, func(role string) bool {
		return role == roleName
	})

	return roleIndex >= 0
}

func (userService) CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error) {

	userEntity := user.User{
		FirstName:     userInput.FirstName,
		LastName:      userInput.LastName,
		Email:         userInput.Email,
		UserName:      userInput.UserName,
		Password:      userInput.Password,
		RegisterDate:  time.Now(),
		CreatorUserId: userInput.CreatorUserId,
	}

	userRepository := repository.NewUserRepository()
	userId, err := userRepository.InsertUser(userEntity)

	return userId, err
}
