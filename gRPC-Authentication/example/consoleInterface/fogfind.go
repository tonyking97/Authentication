package main

import (
	"auth_server/authpb"
	"context"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"gopkg.in/urfave/cli.v2"
	"log"
	"os"
	"regexp"
)

var servStatus = cli.Command{
	Name:"status",
	Usage:"Initiates a gRPC server",
	Action: func(ctx *cli.Context) error {
		var (
			option = grpc.WithInsecure()
		)

		conn, err := grpc.Dial("localhost:8080", option)
		if err != nil {
			fmt.Printf("Did not connect: %v", err)
			return err
		}

		defer conn.Close()
		client := authpb.NewAuthServiceClient(conn)
		res, err := client.CheckServerStatus(context.Background(), &authpb.CheckServerStatusRequest{})
		if err != nil {
			fmt.Printf("Did not connect: %v", err)
			return err
		}

		fmt.Printf("%v\n", res)
		return nil
	},
}

var servAuth = cli.Command{
	Name:"start",
	Usage:"Server Authentication",
	Action: func(ctx *cli.Context) error {
		option := grpc.WithInsecure()

		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial("localhost:8080", option)
		if err != nil {
			log.Fatalf("Did not connect: %v", err)
		}
		defer conn.Close()

		client := authpb.NewAuthServiceClient(conn)
		res := &authpb.CheckUsernameResponse{}
		prompt := promptui.Prompt{
			Label: "Username",
			Validate: func(input string) error {
				reg, _ := regexp.Compile("[a-z]+")
				if len(input) < 5 || !reg.MatchString(input) {
					return errors.New("Username must have at least 5 characters and alphanumeric value")
				}
				return nil
			},

		}
		username, err := prompt.Run()
		res, err =client.CheckUsername(context.Background(), &authpb.CheckUsernameRequest{Username:username})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//d := color.New(color.FgCyan, color.Bold)
		//_, _ = d.Println("Hi " + res.Username + "..!!")


		prompt = promptui.Prompt{
			Label: "Password",
			Mask: 1,
			Validate: func(input string) error {
				if len(input) < 4 {
					return errors.New("password must have at least 4 characters")
				}
				return nil
			},

		}

		password, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loginResponse , err := client.Login(context.Background(),&authpb.LoginRequest{Username:res.Username, Password:password})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println(loginResponse)

		prompt = promptui.Prompt{
			Label: "Device ID",
			Validate: func(input string) error {
				if len(input) < 8 {
					return errors.New("password must have at least 8 characters")
				}
				return nil
			},

		}
		deviceId , err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Name : " , res.Username, "\nAvatar : ", res.Avatar, "\ndeviceId : ", deviceId)
		return nil
	},
}

func main(){
	app := &cli.App{
		Name:"fogfind",
		Usage:"Connecting 'X'",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:"debug",
				Usage:"enables debugging log",
			},
		},
		Commands: []*cli.Command{
			&servStatus,
			&servAuth,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalf("ERROR : %v", err)
	}
}
