package main

import (
	ci "auth_server/example/consoleInterface/console_interface_pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
}

func (*Server) Status(ctx context.Context,req *ci.StatusRequest) (*ci.StatusResponse, error) {
	return &ci.StatusResponse{
		Name:"fogfind server",
		Description:"Connecting X",
		Host:"localhost",
		IpAddress:"localhost",
		Port:"2022",
		Status:"good :)",
	}, nil
}

func main()  {
	var options []grpc.ServerOption
	lis, err := net.Listen("tcp", "localhost:"+strconv.Itoa(2022))
	if err != nil {
		log.Fatalf("could not list on %v: %v", strconv.Itoa(2022), err)
	}

	gRPC:= grpc.NewServer(
		options...
	)
	ci.RegisterConsoleInterfaceServer(gRPC, &Server{})
	err = gRPC.Serve(lis)
	log.Fatalf("Error : %v", err)
}


