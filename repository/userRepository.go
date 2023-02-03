package repository

/*

import (
	"API/database"
	"API/model/user"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	GetUserList() ([]user.User, error)
}

type userRepository struct {
	db database.Db
}

func NewUserRepository() UserRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return userRepository{
		db: db,
	}
}


func (userRepository userRepository) GetUserList() ([]user.User, error) {

	// give us list of users from DB
	userCollection := userRepository.db.GetUserCollection()

	// find instead findone becuase we need a list of users
	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []user.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil

}


*/
