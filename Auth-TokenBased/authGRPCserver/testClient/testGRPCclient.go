package main

import (
	"../../authpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	option := grpc.WithInsecure()

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", option)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	res, err := client.CheckUsername(context.Background(), &authpb.CheckUsernameRequest{Username: "tonyking"})

	log.Println(res, err)
}
