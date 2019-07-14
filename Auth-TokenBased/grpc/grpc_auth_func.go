package grpc

import (
	"../authpb"
	"../core"
	"../models"
	"../validation/govalidator"
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

type server struct {}

//Login
func (*server) Login(ctx context.Context,req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	log.Println("gRPC login request")

	md, _ := metadata.FromIncomingContext(ctx)

	ipDetails := &models.IpDetails{}
	_= json.Unmarshal([]byte(req.Ipdetails), ipDetails)

	requestUser := &models.User{
		Username:req.Username,
		Password:req.Password,
		Localization:req.Localization,
		Ip_details:*ipDetails,
	}


	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &authpb.LoginResponse{
			Success:false,
			Message: err.Error(),
			Token:&authpb.Token{
				Token:"",
			},
		}
		return res,nil
	} else {
		responseStatus, response := core.Login(requestUser, md["user-agent"][0]) //need to pass useragent
		if responseStatus == http.StatusOK {
			res := &authpb.LoginResponse{
				Success:true,
				Message: "success",
				Token:&authpb.Token{
					Token:response,
				},
			}
			return res,nil
		} else {
			res := &authpb.LoginResponse{
				Success:false,
				Message: response,
				Token:&authpb.Token{
					Token:"",
				},
			}
			return res,nil
		}

	}
}

//Check Username
func (*server) CheckUsername(ctx context.Context,req *authpb.CheckUsernameRequest) (*authpb.CheckUsernameResponse, error) {
	log.Println("gRPC check username request")

	requestUser := &models.Username{
		Username:req.Username,
	}

	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &authpb.CheckUsernameResponse{
			Success:false,
			Message:err.Error(),
			Username:"",
			Avatar:"",
		}
		return res, nil
	} else {
		if ack, username, avatar := core.CheckUsername(requestUser); !ack {
			res := &authpb.CheckUsernameResponse{
				Success:false,
				Message:"No such Username exist",
				Username:"",
				Avatar:"",
			}
			return res, nil
		} else {
			res := &authpb.CheckUsernameResponse{
				Success:true,
				Message:"Success",
				Username:username,
				Avatar:avatar,
			}
			return res, nil
		}
	}
}

func (*server) CheckUsernameAvailability(ctx context.Context,req *authpb.CheckUsernameAvailabilityRequest) (*authpb.CheckUsernameAvailabilityResponse, error) {
	log.Println("gRPC check username availability request")

	requestUser := &models.Username{
		Username:req.Username,
	}

	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &authpb.CheckUsernameAvailabilityResponse{
			Success:false,
			Message:err.Error(),
		}
		return res, nil
	} else {
		if ack, _, _ := core.CheckUsername(requestUser); !ack {
			res := &authpb.CheckUsernameAvailabilityResponse{
				Success:true,
				Message:"Username Available",
			}
			return res, nil
		} else {
			res := &authpb.CheckUsernameAvailabilityResponse{
				Success:false,
				Message:"Username already exist. Try another username.",
			}
			return res, nil
		}
	}
}

func (*server) CheckEmailAvailability(ctx context.Context,req *authpb.CheckEmailAvailabilityRequest) (*authpb.CheckEmailAvailabilityResponse, error) {
	log.Println("gRPC check email availability request")

	requestUser := &models.Email{
		Email:req.Email,
	}

	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &authpb.CheckEmailAvailabilityResponse{
			Success:false,
			Message:err.Error(),
		}

		return res, nil
	} else {
		if ack := core.CheckEmail(requestUser); !ack {
			res := &authpb.CheckEmailAvailabilityResponse{
				Success:true,
				Message:"Success",
			}

			return res, nil
		} else {
			res := &authpb.CheckEmailAvailabilityResponse{
				Success:false,
				Message:"EmailID already exists",
			}

			return res, nil
		}
	}
}

func (*server) Signup(ctx context.Context,req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	requestUser := &models.SignUpDetails{
		Email:req.Email,
		Username:req.Username,
		FirstName:req.Firstname,
		LastName:req.Lastname,
		Password:req.Password,
	}

	if _, err := govalidator.CustomValidateStruct(requestUser); err != nil {

		var errArr[] *authpb.SignupError

		for _, element := range err {
			errElem := &authpb.SignupError{
				Name:element.Name,
				Err:element.Err,
			}
			errArr = append(errArr,errElem)
		}

		res := &authpb.SignupResponse{
			Success:false,
			Message:"Validation error",
			Err:errArr,
		}
		return res, nil

	} else {

		if !core.Signup(requestUser) {

			res := &authpb.SignupResponse{
				Success:false,
				Message:"Account creation failed",
			}
			return res, nil
		} else {

			res := &authpb.SignupResponse{
				Success:true,
				Message:"Account created successfully",
			}
			return res, nil
		}
	}
}
