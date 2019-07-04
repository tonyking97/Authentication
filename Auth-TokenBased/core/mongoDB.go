package core

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : File contains all the connection to the MongoDB database and collection.
 */

var client *mongo.Client = nil
var usersCollection *mongo.Collection = nil
var sessionCollection *mongo.Collection = nil

//Initialize MongoDB connection
func MongoDBInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	if client == nil {
		var err error = nil
		client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
}

//Initialize and return the reference of the userList Collection.
func UsersCollectionInit() *mongo.Collection {

	if usersCollection == nil {
		usersCollection = client.Database("users").Collection("userList")

		_, err := usersCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys: bson.M{
					"email": 1,
				},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.M{
					"username": 1,
				},
				Options: options.Index().SetUnique(true),
			},
		})
		if err != nil {
			log.Println(err)
		}
	}

	return usersCollection
}

//Initialize and return the reference of the sessionDetails Collection.
func SessionCollectionInit() *mongo.Collection {

	if sessionCollection == nil {
		sessionCollection = client.Database("users").Collection("sessionDetails")
	}

	return sessionCollection
}
