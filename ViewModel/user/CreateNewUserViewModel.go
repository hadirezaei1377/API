package user

type CreateNewUserViewModel struct {
	LastName      string `validate:"required"`
	FirstName     string
	Email         string
	UserName      string
	Password      string
	CreatorUserId string
}
