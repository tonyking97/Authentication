package main

import (
	"../authpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"strconv"
	"time"
)

func main() {

	option := grpc.WithInsecure()

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", option)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	//checkStatus(client)
	//registerAccount(client)
	loginResponse := login(client)

	var ctx context.Context
	header := metadata.New(map[string]string{"authorization": loginResponse.Token, "time": strconv.Itoa(int(time.Now().UTC().Unix()))})
	ctx = metadata.NewOutgoingContext(context.Background(), header)

	//updateAccount(client,ctx)
	getAccountDetails(client, ctx)
	//changePassword(client,ctx)
	getAllAccounts(client, ctx)
}

func changePassword(client authpb.AuthServiceClient, ctx context.Context) {
	res4, err := client.ChangePassword(ctx, &authpb.ChangePasswordRequest{
		OldPassword:     "nanda",
		NewPassword:     "nanda1",
		ConfirmPassword: "nanda1",
	})

	if err != nil {
		log.Fatalf("Error while calling ChangePassword RPC..!!")
	}
	log.Println("Response from ChangePassword RPC : ", res4)

}

func getAccountDetails(client authpb.AuthServiceClient, ctx context.Context) {
	res3, err := client.GetAccountDetails(ctx, &authpb.GetAccountDetailsRequest{})
	if err != nil {
		log.Fatalf("Error while calling GetAccountDetails RPC..!! %v", err)
	}
	log.Println("Response from GetAccountDetails RPC : ", res3)
}

func updateAccount(client authpb.AuthServiceClient, ctx context.Context) {
	res2, err := client.UpdateAccount(ctx, &authpb.UpdateAccountRequest{
		Username:  "nanda",
		Email:     "nandakumar111@outlook.com",
		FirstName: "Nandakumar",
		LastName:  "R",
	})

	if err != nil {
		log.Fatalf("Error while calling UpdateAccount RPC..!!")
	}
	log.Println("Response from UpdateAccount RPC : ", res2)
}

func login(client authpb.AuthServiceClient) *authpb.LoginResponse {
	res1, err := client.Login(context.Background(), &authpb.LoginRequest{Username: "nanda", Password: "nanda"})
	if err != nil {
		log.Fatalf("Error while calling Login RPC..!! %v", err)
	}
	log.Println("Response from Login RPC : ", res1)
	return res1
}

func checkStatus(client authpb.AuthServiceClient) {
	res, err := client.CheckServerStatus(context.Background(), &authpb.CheckServerStatusRequest{})
	if err != nil {
		log.Fatalf("Error while calling CheckServerStatus RPC..!!")
	}
	log.Println("Response from CheckServerStatus RPC : ", res)

}

func registerAccount(client authpb.AuthServiceClient) {
	res, err := client.RegisterAccount(context.Background(), &authpb.RegisterAccountRequest{
		Username:        "nanda",
		Email:           "nandakumar111@outlook.com",
		Password:        "nanda",
		ConfirmPassword: "nanda",
		FirstName:       "Nandakumar",
		LastName:        "R",
		CompanyTerms:    true,
		GovtTerms:       true,
	})

	if err != nil {
		log.Fatalf("Error while calling RegisterAccount RPC..!!")
	}
	log.Println("Response from RegisterAccount RPC : ", res)

}

func getAllAccounts(client authpb.AuthServiceClient, ctx context.Context) {
	log.Println("Staring to do server streaming RPC..!!")
	req := &authpb.GetAllAccountsRequest{
		Skip:  0,
		Limit: 10,
	}

	resStream, err := client.GetAllAccounts(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling GetAllAccounts RPC : %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming result from server : %v", err)
		}
		log.Printf("Response from GetAllAccounts : %v", res)
	}
}
