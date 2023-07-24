package user

type CreateNewUserViewModel struct {
	LastName      string `validate:"required"`
	FirstName     string `validate:"required"`
	Email         string `validate:"required"`
	UserName      string `validate:"required"`
	Password      string `validate:"required"`
	CreatorUserId string `validate:"required"`
}
type EditUserViewModel struct {
	TargetUserId string
	LastName     string `validate:"required"`
	FirstName    string `validate:"required"`
	Email        string `validate:"required"`
	UserName     string `validate:"required"`
	Password     string `validate:"required"`
}
