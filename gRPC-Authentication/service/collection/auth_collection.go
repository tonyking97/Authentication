package collection

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

var Client *mongo.Client = nil
var AccountCollection *mongo.Collection = nil
var SessionCollection *mongo.Collection = nil

//Initialize MongoDB connection
func MongoDBInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	if Client == nil {
		var err error = nil
		Client , err = mongo.Connect(context.TODO(), clientOptions)
		if err!= nil{
			log.Fatal(err)
		}
	}

	err := Client.Ping(context.TODO(), nil)

	if err!= nil{
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB :) ")
}

func AccountCollectionInit() *mongo.Collection{
	if AccountCollection == nil{
		AccountCollection = Client.Database("authentication").Collection("account")

		_, err := AccountCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys:bson.M{
					"email":1,
				},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys:bson.M{
					"username":1,
				},
				Options: options.Index().SetUnique(true),
			},
		})
		if err!= nil{
			log.Println(err)
		}
	}

	return AccountCollection
}

func SessionCollectionInit()  *mongo.Collection{
	if SessionCollection == nil{
		SessionCollection = Client.Database("authentication").Collection("session")

		_, err := SessionCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys:bson.M{
					"browser_details.app_name":1,
				},
				Options: options.Index().SetUnique(true),
			},
		})
		if err!= nil{
			log.Println(err)
		}
	}

	return SessionCollection
}
