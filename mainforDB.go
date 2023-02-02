package main

// methods for connecting to DB and recieve data from it

/*

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	userCollection := client.Database("toplearn-api-golang").Collection("users")

	var res bson.M
	err = userCollection.FindOne(ctx, bson.D{{"firstName", "Mohammad"}}).Decode(&res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

}


*/
