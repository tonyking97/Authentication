package main

import (
	"auth_server/example/file_transfer/pb"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"io"
	"os"
	"strconv"
)

const (
	address     = "localhost:50051"
	sentValue = 1000000 //limit
)

func main()  {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("didn't connect %v", err)
	}
	defer conn.Close()
	c := pb.NewGuploadServiceClient(conn)

	var (
		buf     []byte
		status  *pb.UploadStatus
	)
	// open input file
	fi, err := os.Open("./input.dmg")
	if err != nil {
		fmt.Println("Not able to open")
		return
	}

	stat, err := fi.Stat()
	if err != nil {
		return
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			fmt.Println("Not able to open")
			return
		}
	}()

	ctx := context.Background()
	stream, err := c.Upload(ctx)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create upload stream for file %s",
			fi)
		return
	}
	defer stream.CloseSend()

	buf = make([]byte,stat.Size())
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			err = errors.Wrapf(err,
				"failed to send chunk via stream")
			return
		}
		if n == 0 {
			break
		}
		var i int64
		for i = 0 ; i < ((stat.Size()/sentValue)*sentValue)  ; i += sentValue {
			err = stream.Send(&pb.Chunk{
				Content: buf[i:i+sentValue],
				TotalSize:strconv.FormatInt(stat.Size(), 10),
				Received:strconv.FormatInt(i+sentValue, 10),
			})
		}
		if stat.Size()%sentValue > 0{
			err = stream.Send(&pb.Chunk{
				Content: buf[((stat.Size()/sentValue)*sentValue):((stat.Size()/sentValue*sentValue)+ (stat.Size()%sentValue))],
				TotalSize:strconv.FormatInt(stat.Size(), 10),
				Received:string(stat.Size()%sentValue),
			})
		}

		if err != nil {
			err = errors.Wrapf(err,
				"failed to send chunk via stream")
			return
		}
	}


	status, err = stream.CloseAndRecv()
	if err != nil {
		err = errors.Wrapf(err,
			"failed to receive upstream status response")
		return
	}

	if status.Code != pb.UploadStatusCode_Ok {
		err = errors.Errorf(
			"upload failed - msg: %s",
			status.Message)
		return
	}

	return
}
