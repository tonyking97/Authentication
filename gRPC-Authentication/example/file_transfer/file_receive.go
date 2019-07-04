package main

import (
	"./pb"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

const (
	port = "localhost:50051"
)

type server struct {
}

func (s *server) Upload(stream pb.GuploadService_UploadServer) (err error) {
	// open output file
	fo, err := os.Create("./output")
	if err != nil {
		return errors.New("failed to create file")
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	var res *pb.Chunk
	for {
		res, err = stream.Recv()

		if err == io.EOF {
			err = stream.SendAndClose(&pb.UploadStatus{
				Message: "Upload received with success",
				Code:    pb.UploadStatusCode_Ok,
			})
			if err != nil {
				err =  errors.New("failed to send status code")
				return err
			}
			return nil
		}
		fmt.Println(res.Received)
		// write a chunk
		if _, err := fo.Write(res.Content); err != nil {
			err =  errors.New(err.Error())
			return err
		}
	}
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterGuploadServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
