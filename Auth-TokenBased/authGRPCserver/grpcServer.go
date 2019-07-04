package main

import (
	"../authpb"
	"../core"
	"./services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	core.MongoDBInit()

	// create a listener on TCP port
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could not list on localhost:50051 %v", err)
	}

	// Create an empty array of gRPC options
	var options []grpc.ServerOption

	// create a gRPC server object
	gRPC := grpc.NewServer(
		options...,
	)

	authpb.RegisterAuthServiceServer(gRPC, &services.AuthGRPCServer{})
	log.Printf("Starting HTTP/2 gRPC Authentication Server on %s : %s", "127.0.0.1", "50051")

	// start the server
	log.Println(gRPC.Serve(lis))
}
