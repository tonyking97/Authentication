package grpc

import (
	"../authpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartServer() error {
	lis, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen",err)
	}

	s:=grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &server{})

	reflection.Register(s)

	log.Println("Starting gRPC server in port 50051")

	if err:= s.Serve(lis); err!=nil {
		log.Fatalln("Failed to serve",err)
	}

	return nil;
}