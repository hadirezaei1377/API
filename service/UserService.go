package service

// service for improrve the Project architecture

import "API/model/user"

// an user can user interface for using interface but not the main userservice

type UserService interface {
	GetUserList() ([]user.User, error)
}

// private struct for main userservice
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
