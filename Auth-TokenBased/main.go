package main

import (
	"./core"
	"./grpc"
	"./routers"
	"github.com/mileusna/crontab"
	"github.com/rs/cors"
	"log"
	"net/http"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Main function of the Authentication Service
 */

func main() {
	//Initializing and Establishing connection with MongoDB
	core.MongoDBInit()

	//Initializing router to router the incoming request URL
	router := routers.InitRouters()

	//Initializing cron job table
	ctab := crontab.New()

	//Adding Token check Job to Cron job table
	ctab.MustAddJob("* * * * *", core.CheckTokenStatus) //Check Expiry time of token every minute and change the status

	//Enabling CORS. Caution : Only need for dev Environment. Turn off CORS in production env by making Secure Connection.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or we can have our header key values which we are using in our application
		},
	})

	//Starting gRPC server
	go func() {
		err := grpc.StartServer()
		if err != nil {
			log.Fatalf("Failed to start gRPC server: %s", err)
		}
	}()

	//Starting the Authentication Server
	handler := c.Handler(router)
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Server creation failed.")
	}
}
