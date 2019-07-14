package grpc

import (
	"../authDashboardpb"
	"context"
)

type serverdashboard struct {}

func (*serverdashboard) CheckToken(ctx context.Context,req *authdashboardpb.CheckTokenRequest) (*authdashboardpb.CheckTokenResponse, error) {
	return &authdashboardpb.CheckTokenResponse{
		Success:true,
		Message:"token valid",
	} , nil
}

