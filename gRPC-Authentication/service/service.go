package service

import (
	"../authpb"
	"../service/core_function"
	"../service/models"
	"../service/service_function"
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type AuthServer struct {
}

func (*AuthServer) CheckServerStatus(ctx context.Context, req *authpb.CheckServerStatusRequest) (*authpb.CheckServerStatusResponse, error) {
	log.Println("CheckServerStatus RPC Initiated..!!")
	return &authpb.CheckServerStatusResponse{
		ServerName: "Authentication Server",
		Host:       "Localhost",
		Port:       "50051",
		Time:       time.Now().Format(time.Stamp),
		Status:     "Good :)",
	}, nil
}

func (*AuthServer) RegisterAccount(ctx context.Context, req *authpb.RegisterAccountRequest) (*authpb.RegisterAccountResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, status.Errorf(codes.InvalidArgument, "Password and confirm password are not same..!!")
	}
	uuid, _ := core_function.NewV4()
	hashes := md5.New()
	hashes.Write([]byte(req.Password))
	reqTime := time.Now().UTC().Unix()

	model := &models.AccountModel{
		Id:          uuid.String(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Username:    req.Username,
		Password:    hex.EncodeToString(hashes.Sum(nil)),
		Email:       req.Email,
		CreatedTime: reqTime,
		UpdatedTime: reqTime,
		Status:      "pending",
	}

	log.Println("RegisterAccount invoked with : ", model)
	err := service_function.RegisterAccountFunction(model)
	if err != nil {
		return nil, err
	}
	return &authpb.RegisterAccountResponse{Result: "Account created successfully..!!", Status: "pending"}, nil
}

func (*AuthServer) CheckUsername(ctx context.Context, req *authpb.CheckUsernameRequest) (*authpb.CheckUsernameResponse, error) {
	return service_function.CheckUsernameFunction(req.Username)
}

func (*AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	hashes := md5.New()
	hashes.Write([]byte(req.Password))
	req.Password = hex.EncodeToString(hashes.Sum(nil))

	log.Println("Login invoked with : ", req)

	res, err := service_function.LoginAccountFunction(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (*AuthServer) UpdateAccount(ctx context.Context, req *authpb.UpdateAccountRequest) (*authpb.UpdateAccountResponse, error) {
	log.Println("Login invoked with : ", req)
	return &authpb.UpdateAccountResponse{Result: "Account updated successfully..!!"}, nil
}

func (*AuthServer) GetAccountDetails(ctx context.Context, req *authpb.GetAccountDetailsRequest) (*authpb.GetAccountDetailsResponse, error) {
	log.Println("GetAccountDetails invoked with : ", req)
	meta, _ := metadata.FromIncomingContext(ctx)
	fs_id, _ := core_function.DecryptString(meta["authorization"][0])
	accountDetails, err := service_function.GetAccountDetailsFunction(fs_id)
	if err != nil {
		return nil, err
	}

	return &authpb.GetAccountDetailsResponse{
		FirstName: accountDetails.FirstName,
		LastName:  accountDetails.LastName,
		Username:  accountDetails.Username,
		Email:     accountDetails.Email,
	}, nil
}

func (*AuthServer) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	log.Println("ChangePassword invoked with : ", req)
	return &authpb.ChangePasswordResponse{}, nil
}

func (*AuthServer) GetAllAccounts(req *authpb.GetAllAccountsRequest, stream authpb.AuthService_GetAllAccountsServer) error {
	log.Printf("GetAccountDetails function invoked  with %v", stream.Context())
	allAccounts := service_function.GetAllAccountsFunction(req)
	for _, account := range allAccounts {
		res := &authpb.GetAllAccountsResponse{
			FfId:        account.Id,
			FirstName:   account.FirstName,
			LastName:    account.LastName,
			Email:       account.Email,
			Username:    account.Username,
			CreatedTime: account.CreatedTime,
			Status:      account.Status,
		}
		_ = stream.Send(res)
	}
	return nil
}

func (*AuthServer) GetAllSessionDetails(req *authpb.GetAllSessionDetailsRequest, stream authpb.AuthService_GetAllSessionDetailsServer) error {
	log.Printf("GetAllAccountSessionDetails function invoked  with %v", req)
	sessionDetails := service_function.GetAllSessionDetailsFunction(req.FfId)
	for _, session := range sessionDetails {
		log.Println(session.Id)
		res := &authpb.GetAllSessionDetailsResponse{
			FsId:          session.Id,
			LoggedIn:      string(session.ExpiryTime),
			CurrentStatus: session.Status,
			LastPing:      string(time.Now().Nanosecond()),
		}
		_ = stream.Send(res)
	}
	return nil
}
