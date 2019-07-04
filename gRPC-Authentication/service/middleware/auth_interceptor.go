package middleware

import (
	"../../service/core_function"
	"../../service/service_function"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

// Interceptor for auth check
func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, _ := metadata.FromIncomingContext(ctx)
	log.Printf("MetaData : %v", meta)
	log.Println(info.FullMethod)
	withOutAuthMethods := []string{"/authpb.AuthService/CheckServerStatus", "/authpb.AuthService/RegisterAccount", "/authpb.AuthService/CheckUsername", "/authpb.AuthService/Login"}

	isThere, _ := core_function.InArray(info.FullMethod, withOutAuthMethods)
	if !isThere {
		if meta["authorization"][0] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Authorization token not found..!!")
		}
		fsId, _ := core_function.DecryptString(meta["authorization"][0])
		sessionDetails, err := service_function.AuthCheck(fsId)
		if err != nil {
			return nil, err
		}
		newCtx := context.WithValue(ctx, "session_details", sessionDetails)
		return handler(newCtx, req)
	}
	return handler(ctx, req)
}

func AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("Method : " + info.FullMethod)
	log.Printf("ctx : %v", ss.Context())
	withOutAuthMethods := []string{"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"}
	meta, _ := metadata.FromIncomingContext(ss.Context())
	isThere, _ := core_function.InArray(info.FullMethod, withOutAuthMethods)
	if !isThere {
		ffId, _ := core_function.DecryptString(meta["authorization"][0])
		sessionDetails, err := service_function.AuthCheck(ffId)
		if err != nil {
			return err
		}
		_ = context.WithValue(ss.Context(), "session_details", sessionDetails)
		return handler(srv, ss)
	}
	return handler(srv, ss)
}
