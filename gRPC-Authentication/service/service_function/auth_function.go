package service_function

import (
	"../../authpb"
	authcollection "../../service/collection"
	"../../service/core_function"
	"../../service/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func RegisterAccountFunction(req *models.AccountModel) error {
	log.Printf("SignUpAccountFunction invoked with : %v", req)
	accountCollection := authcollection.AccountCollectionInit()
	res, err := accountCollection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Printf("Error : %v", err)
		return status.Errorf(codes.Canceled, fmt.Sprintf("Account failed to create. Reason is %v", err))
	}
	log.Printf("Account Created : %v", res)
	return nil
}

func CheckUsernameFunction(username string) (*authpb.CheckUsernameResponse, error) {
	accountCollection := authcollection.AccountCollectionInit()
	account := models.AccountModel{}
	err := accountCollection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&account)
	if err != nil {
		log.Printf("Error : %v", err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Record not found..!!"))
	}
	firstName := account.FirstName
	lastName := account.LastName
	avatar := string(firstName[0]) + string(lastName[0])
	return &authpb.CheckUsernameResponse{Username: username, Avatar: avatar}, nil
}

func LoginAccountFunction(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	log.Printf("LoginAccountFunction invoked with : %v", req)

	account, err := VerifyUser(req)
	if err != nil {
		return nil, err
	}

	log.Println("account : ", account)
	sessionDetails, err := CheckSession(ctx, account)
	if err != nil {
		return nil, err
	}
	encrypt, _ := core_function.EncryptString(sessionDetails.Id)
	return &authpb.LoginResponse{Token: encrypt}, nil
}

func VerifyUser(req *authpb.LoginRequest) (*models.AccountModel, error) {
	accountCollection := authcollection.AccountCollectionInit()
	account := models.AccountModel{}
	err := accountCollection.FindOne(context.TODO(), bson.D{{"username", req.Username}}).Decode(&account)
	if err != nil {
		log.Printf("Error : %v", err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Record not found..!!"))
	} else if account.Password != req.Password {
		log.Printf("Password Incorrect :(")
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Username or Password Incorrect..!!"))
	}

	if account.Status != "active" {
		return nil, status.Errorf(codes.PermissionDenied, fmt.Sprintf("Account is in "+account.Status+" state..!! Please contact administrator."))
	}
	return &account, nil
}

func CheckSession(ctx context.Context, account *models.AccountModel) (*models.SessionDetails, error) {
	meta, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Metadata : %v", meta)
	sessionCollection := authcollection.SessionCollectionInit()
	var sessionDetails []*models.SessionDetails
	session := models.SessionDetails{}
	cur, _ := sessionCollection.Find(context.TODO(), bson.D{{"user_details.username", account.Username}}, options.Find())
	for cur.Next(context.TODO()) {
		var elem models.SessionDetails
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		sessionDetails = append(sessionDetails, &elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	_ = cur.Close(context.TODO())

	if len(sessionDetails) == 0 {
		uuid, _ := core_function.NewV4()
		reqTime := time.Now().UTC().Unix()
		expTime := time.Now().Add(time.Hour).UTC().Unix()
		session = models.SessionDetails{
			Id:   uuid.String(),
			FfId: account.Id,
			UserDetails: models.UserDetails{
				Username:  account.Username,
				Email:     account.Email,
				FirstName: account.FirstName,
				LastName:  account.LastName,
			},
			LoggedTime: reqTime,
			ExpiryTime: expTime,
		}
		_, err := sessionCollection.InsertOne(context.TODO(), session)
		if err != nil {
			log.Printf("insert error :%v", err)
			return nil, err
		}
		return &session, nil
	}
	return sessionDetails[0], nil
}

func AuthCheck(fsId string) (*models.SessionDetails, error) {
	sessionCollection := authcollection.SessionCollectionInit()
	sessionDetails := models.SessionDetails{}
	err := sessionCollection.FindOne(context.TODO(), bson.D{{"_id", fsId}}).Decode(&sessionDetails)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Channel closed..!!")
	}
	return &sessionDetails, err
}

func GetAllAccountsFunction(req *authpb.GetAllAccountsRequest) []*models.AccountModel {
	accountCollection := authcollection.AccountCollectionInit()
	var allAccounts []*models.AccountModel
	findOptions := options.Find()
	findOptions.SetLimit(req.Limit).SetSkip(req.Skip)
	cur, _ := accountCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem models.AccountModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil
		}
		allAccounts = append(allAccounts, &elem)
	}
	if err := cur.Err(); err != nil {
		return nil
	}
	_ = cur.Close(context.TODO())
	return allAccounts
}

func GetAccountDetailsFunction(fsId string) (*models.AccountModel, error) {
	sessionCollection := authcollection.SessionCollectionInit()
	sessionDetails := models.SessionDetails{}
	err := sessionCollection.FindOne(context.TODO(), bson.D{{"_id", fsId}}).Decode(&sessionDetails)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Session record not Found ..!!"))
	}
	accountCollection := authcollection.AccountCollectionInit()
	var accountDetails models.AccountModel
	err = accountCollection.FindOne(context.TODO(), bson.D{{"_id", sessionDetails.FfId}}).Decode(&accountDetails)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Account record not Found ..!! %v", err))
	}
	return &accountDetails, nil
}

func GetAllSessionDetailsFunction(ff_id string) []*models.SessionDetails {
	sessionCollection := authcollection.SessionCollectionInit()
	var sessionDetails []*models.SessionDetails
	cur, _ := sessionCollection.Find(context.TODO(), bson.D{{"ff_id", ff_id}})
	for cur.Next(context.TODO()) {
		var elem models.SessionDetails
		err := cur.Decode(&elem)
		if err != nil {
			return nil
		}
		sessionDetails = append(sessionDetails, &elem)
	}
	if err := cur.Err(); err != nil {
		return nil
	}
	_ = cur.Close(context.TODO())
	return sessionDetails
}

func AccountCreateFunction() {

}

func GetAccountFunction() {

}

func UpdateAccountFunction() {

}

func DeleteAccountFunction() {

}
