package grpc

import (
	"../authDashboardpb"
	"../core"
	"../models"
	"context"
)

type authdashboard struct {}

func (*authdashboard) CheckToken(ctx context.Context,req *authdashboardpb.CheckTokenRequest) (*authdashboardpb.CheckTokenResponse, error) {

	return &authdashboardpb.CheckTokenResponse{
		Success:true,
		Message:"token valid",
	} , nil
}

func (*authdashboard) UserDetails(ctx context.Context,req *authdashboardpb.UserDetailsRequest) (*authdashboardpb.UserDetailsResponse, error) {

	if ack, result := core.GetUserDetails(&models.Ff_id_user{Ff_id: ctx.Value("ff_id").(string)}); !ack {
		res := &authdashboardpb.UserDetailsResponse{
			Success:false,
		}
		return res, nil
	} else {
		res := &authdashboardpb.UserDetailsResponse{
			Success:true,
			Username:result.Username,
			Firstname:result.Firstname,
			Lastname:result.Lastname,
			Email:result.Email,
		}
		return res,nil
	}
}

func (*authdashboard) CurrentSessionDetails(ctx context.Context,req *authdashboardpb.CurrentSessionDetailsRequest) (*authdashboardpb.CurrentSessionDetailsResponse, error) {
	ff_id := ctx.Value("ff_id").(string)
	fs_id := ctx.Value("fs_id").(string)

	result := core.GetCurrentSessionList(fs_id)
	tokenCount := core.GetActiveTokensCount(ff_id)
	sessionCount := core.GetActiveSessionsCount(ff_id)

	res := &authdashboardpb.CurrentSessionDetailsResponse{
		Activesessions:sessionCount,
		Activetokens:tokenCount,
		Ip:result.Ip_Details.Ip,
		City:result.Ip_Details.City,
		Latitude:result.Ip_Details.Latitude,
		Longitude:result.Ip_Details.Longitude,
	}
	return res,nil
}

func (*authdashboard) SessionDetails(ctx context.Context,req *authdashboardpb.SessionDetailsRequest) (*authdashboardpb.SessionDetailsResponse, error) {
	ff_id := ctx.Value("ff_id").(string)
	results := core.GetSessionList(ff_id)

	var sessionDetails[] *authdashboardpb.SessionDetailLimited

	for _, element := range results {
		sessionDetailsElem := &authdashboardpb.SessionDetailLimited{
			Token:element.Token,
			Status:element.Status,
			LoggedTime:element.LoggedTime,
			LoggedTimeUTC:element.LoggedTimeUTC,
			ExpiryTime:element.ExpiryTime,
			ExpiryTimeUTC:element.ExpiryTimeUTC,
			Mobile:element.Mobile,
			BrowserName:element.BrowserName,
			BrowserVersion:element.BrowserVersion,
			OsName:element.OsName,
			OsPlatform:element.Platform,
		}
		sessionDetails = append(sessionDetails, sessionDetailsElem)
	}

	res := &authdashboardpb.SessionDetailsResponse{
		SessionDetails:sessionDetails,
	}

	return res, nil
}

func (*authdashboard) Logout(ctx context.Context,req *authdashboardpb.LogoutRequest) (*authdashboardpb.LogoutResponse, error) {
	ff_id := ctx.Value("ff_id").(string)
	fs_id := ctx.Value("fs_id").(string)

	err := core.Logout(ff_id,fs_id)

	if err!=nil {
		res:= &authdashboardpb.LogoutResponse{
			Success:false,
			Message:err.Error(),
		}
		return res, nil
	} else {
		res:= &authdashboardpb.LogoutResponse{
			Success:true,
		}
		return res, nil
	}
}

