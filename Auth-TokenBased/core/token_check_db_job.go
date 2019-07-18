package core

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Contains method which checks the token expiry time and change the status to inactive while time expires.
 */

type sessionDetails struct{
	FS_Id      string    `bson:"fs_id"`
	ExpiryTime int64     `bson:"expiry_Time"`
}

//Check Expiry time of token every minute and change the status
func CheckTokenStatus() {

	count := 0
	sessionCollection = SessionCollectionInit()

	filter := bson.D{
		{"status", "active"},
	}

	cur, err := sessionCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem sessionDetails
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}
		if elem.ExpiryTime < time.Now().Unix() {
			filter = bson.D{
				{"fs_id", elem.FS_Id},
			}

			update := bson.D{
				{"$set", bson.D{
					{"status", "inactive"},
				}},
			}

			_, err := sessionCollection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				log.Println(err)
			} else {
				count++
			}

		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	log.Println("Cron Job done. Token status changed for ", count, " Documents.")

}
