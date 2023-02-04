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

// read from db
// read section of CRUD operation

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

// Read section of CRUd related to db
/*

func (userRepository userRepository) GetUserById(id string) (user.User, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user.User{}, err
	}
	userCollection := userRepository.db.GetUserCollection()
	var userObject user.User
	// read codes from MongoDB documantation
	// there is BsonD, BsonA, BsonE, ...
	// its about orders,
    // When using it, it asks if the order is important to you or not
    // we have document so we use bsonD
    // The order is not important for us and we just want our ID to be the same as the ID in the database.

	// db.getCollection('users').find({"_id" : ObjectId("6297bdcc8d7757574658ed66")})

	// _id is key for map
	err = userCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectId},
	}).Decode(&userObject)

	if err != nil {
		return user.User{}, err
	}

	return userObject, nil

}
*/
