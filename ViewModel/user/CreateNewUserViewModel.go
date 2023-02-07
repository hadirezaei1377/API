package user

type CreateNewUserViewModel struct {
	LastName  string
	FirstName string
	Email     string
	UserName  string
	Password  string
}

/*

view model used for works like changing password by user
view model full this struct and send it to serverside
we dont need newpassword field becuase this is a field for client (just check) and is not valuable for storge
we use this struct in controller

*/
