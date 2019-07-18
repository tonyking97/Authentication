package grpc

import (
	"../authDashboardpb"
	"../authpb"
	"../grpcMiddleware"
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

	var options []grpc.ServerOption

	//Interceptor
	options = append(options, grpc.UnaryInterceptor(grpcMiddleware.AuthDashboardUnaryInterceptor))

	s:=grpc.NewServer( options... )

	authpb.RegisterAuthServiceServer(s, &server{})
	authdashboardpb.RegisterAuthDashboardServiceServer(s,&authdashboard{})

	reflection.Register(s)

	log.Println("Starting gRPC server in port 50051")

	if err:= s.Serve(lis); err!=nil {
		log.Fatalln("Failed to serve",err)
	}

	return nil;
}