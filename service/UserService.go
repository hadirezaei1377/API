package service

import "API/model/user"

type UserService interface {
	GetUserList() ([]user.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (userService) GetUserList() ([]user.User, error) {
	userList := []user.User{
		{
			FirstName:   "Mohammad",
			LastName:    "Ghari",
			Age:         27,
			PhoneNumber: "0912345678",
		},
		{
			FirstName:   "Ali",
			LastName:    "Rezae",
			Age:         26,
			PhoneNumber: "0913345678",
		},
		{
			FirstName:   "Iman",
			LastName:    "Madaeni",
			Age:         32,
			PhoneNumber: "0914345678",
		},
	}

	return userList, nil
}
