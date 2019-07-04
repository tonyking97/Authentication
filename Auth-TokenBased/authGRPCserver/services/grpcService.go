package services

import (
	"../../authpb"
	"../../core"
	"../../models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGRPCServer struct {
}

func (*AuthGRPCServer) CheckUsername(cxt context.Context, req *authpb.CheckUsernameRequest) (*authpb.CheckUsernameResponse, error) {
	if ack, username, avatar := core.CheckUsername(&models.Username{req.Username}); !ack {
		return nil, status.Errorf(codes.NotFound, "No such Username exist")
	} else {
		return &authpb.CheckUsernameResponse{Username: username, Avatar: avatar}, nil
	}
}
