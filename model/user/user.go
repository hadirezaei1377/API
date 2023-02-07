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

import (
	"time"
)

type User struct {
	Id           string    `bson:"_id,omitempty"`
	FirstName    string    `bson:"firstName,omitempty"`
	LastName     string    `bson:"lastName,omitempty"`
	Email        string    `bson:"Email,omitempty"`
	UserName     string    `bson:"UserName,omitempty"`
	Password     string    `bson:"Password,omitempty"`
	RegisterDate time.Time `bson:"RegisterDate,omitempty"`
}
