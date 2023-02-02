package database

/*

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// this struct used for access from out of this package to database

type Db struct {
	client *mongo.Client
}

func Connect() (Db, error) {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Db{}, err
	}

	// ping method

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return Db{}, err
	}

	return Db{client: client}, nil

	//
	//userCollection := client.Database("toplearn-api-golang").Collection("users")
	//
	//var res bson.M
	//err = userCollection.FindOne(ctx, bson.D{{"firstName", "Mohammad"}}).Decode(&res)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(res)
}

// a function for using private client in db struct
// its private because we want use it just in db package

func (db Db) GetUserCollection() *mongo.Collection {
	userCollection := db.client.Database("toplearn-api-golang").Collection("users")

	return userCollection
}


*/
