package user

// normal mod
/*

// related struct to user
type User struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string
}

*/

// related to db connection
// db uses bson

type User struct {
	Id          string `bson:"_id,omitempty"`
	FirstName   string `bson:"firstName,omitempty"`
	LastName    string `bson:"lastName,omitempty"`
	Age         int    `bson:"age,omitempty"`
	PhoneNumber string `bson:"phonenumber,omitempty"`
}
