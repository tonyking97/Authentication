package main

import (
	"./config"
	"./network_method"
	"./service/collection"
	"encoding/json"
	"log"
	"os"
)

const configFile = "config/config.json"

func main() {

	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	serverConfig := config.ServerConfig{}
	err := decoder.Decode(&serverConfig)
	if err != nil {
		log.Fatalf("Error while decoding certificates : %v", err)
	}

	collection.MongoDBInit()
	// fire the gRPC server in a goroutine
	go func() {
		err := network_method.StartGRPCServer(serverConfig.Server, serverConfig.TLS)
		if err != nil {
			log.Fatalf("Failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop..!!")
	select {}
}
