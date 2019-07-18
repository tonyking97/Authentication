package grpcMiddleware

import (
	"../core"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"regexp"
	"time"
)

var authPattern,_ = regexp.Compile(`^/[a-z]*[.]`)

// Interceptor for Auth Dashboard Unary call
func AuthDashboardUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()

	if !authCall(info.FullMethod) {
		md, _ := metadata.FromIncomingContext(ctx)
		token := md["token"]

		if len(token) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "token not found")
		}

		ff_id, fs_id, err := core.CheckToken(token[0])
		ctx = context.WithValue(ctx, "ff_id",ff_id)
		ctx = context.WithValue(ctx, "fs_id",fs_id)
		if err != nil {
			return nil, err
		}
	}

	h, err := handler(ctx, req)
	log.Printf("Request - Method:%s\tDuration:%s\n",info.FullMethod,time.Since(start))
	return h,err
}

func authCall(call string) bool {

	if authPattern.FindString(call) == "/authpb." {
		return true
	} else {
		return false
	}
}