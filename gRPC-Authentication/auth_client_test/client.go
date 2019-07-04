package main

import (
	"../authpb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
)


func main()  {

	option := grpc.WithInsecure()

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:8080", option)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	//checkStatus(client)
	_ = receiveFile(client)

	//systemDetails(client)
	//uploadFile(client, )
	//registerAccount(client)
	//loginResponse := login(client)

	//var ctx context.Context
	//header := metadata.New(map[string]string{"authorization":loginResponse.Token, "time": strconv.Itoa(int(time.Now().UTC().Unix()))})
	//ctx = metadata.NewOutgoingContext(context.Background(), header)


	//updateAccount(client,ctx)
	//getAccountDetails(client,ctx)
	//changePassword(client,ctx)
	//getAllAccounts(client,ctx)
}

func receiveFile(client authpb.AuthServiceClient)  error{
	// open output file
	fo, err := os.Create("./README.md")
	if err != nil {
		log.Fatalln(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	resStream, err := client.SendBuildFileToClient(context.Background(), &authpb.UploadRequest{})
	for {
		res, errStream := resStream.Recv()
		if errStream == io.EOF {
			return nil
		}
		// write a chunk
		if _, err := fo.Write(res.Content); err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

	}
}

func changePassword(client authpb.AuthServiceClient, ctx context.Context) {
	res4, err := client.ChangePassword(ctx, &authpb.ChangePasswordRequest{
		OldPassword:"nanda",
		NewPassword:"nanda1",
		ConfirmPassword:"nanda1",
	})

	if err != nil{
		log.Fatalf("Error while calling ChangePassword RPC..!!")
	}
	log.Println("Response from ChangePassword RPC : ", res4)

}

type ClientGRPC struct {
	conn      *grpc.ClientConn
	chunkSize int
}


func getAccountDetails(client authpb.AuthServiceClient, ctx context.Context) {
	res3, err := client.GetAccountDetails(ctx, &authpb.GetAccountDetailsRequest{})
	if err != nil{
		log.Fatalf("Error while calling GetAccountDetails RPC..!! %v",err)
	}
	log.Println("Response from GetAccountDetails RPC : ", res3)
}

func updateAccount(client authpb.AuthServiceClient, ctx context.Context) {
	res2, err := client.UpdateAccount(ctx, &authpb.UpdateAccountRequest{
		Username:"nanda",
		Email:"nandakumar111@outlook.com",
		FirstName:"Nandakumar",
		LastName:"R",
	})

	if err != nil{
		log.Fatalf("Error while calling UpdateAccount RPC..!!")
	}
	log.Println("Response from UpdateAccount RPC : ", res2)
}

func login(client authpb.AuthServiceClient) *authpb.LoginResponse {
	res1, err := client.Login(context.Background(), &authpb.LoginRequest{Username:"nanda", Password:"nanda"})
	if err != nil{
		log.Fatalf("Error while calling Login RPC..!! %v",err)
	}
	log.Println("Response from Login RPC : ", res1)
	return res1
}

func checkStatus(client authpb.AuthServiceClient) {
	res, err := client.CheckServerStatus(context.Background(), &authpb.CheckServerStatusRequest{})
	if err != nil{
		log.Fatalf("Error while calling CheckServerStatus RPC..!!")
	}
	log.Println("Response from CheckServerStatus RPC : ", res)

}

func registerAccount(client authpb.AuthServiceClient)  {
	res, err := client.RegisterAccount(context.Background(), &authpb.RegisterAccountRequest{
		Username:"nanda",
		Email:"nandakumar111@outlook.com",
		Password:"nanda",
		ConfirmPassword:"nanda",
		FirstName:"Nandakumar",
		LastName:"R",
		CompanyTerms:true,
		GovtTerms:true,
	})

	if err != nil{
		log.Fatalf("Error while calling RegisterAccount RPC..!!")
	}
	log.Println("Response from RegisterAccount RPC : ", res)

}


func getAllAccounts(client authpb.AuthServiceClient, ctx context.Context)  {
	log.Println("Staring to do server streaming RPC..!!")
	req := &authpb.GetAllAccountsRequest{
		Skip:0,
		Limit:10,
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

func systemDetails(client authpb.AuthServiceClient){
	//----------------------
	// Get the local machine IP address
	// https://www.socketloop.com/tutorials/golang-how-do-I-get-the-local-ip-non-loopback-address
	//----------------------

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}

	var currentIP, currentNetworkHardwareName string

	for _, address := range addrs {

		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("Current IP address : ", ipnet.IP.String())
				currentIP = ipnet.IP.String()
			}
		}
	}

	//fmt.Println("------------------------------")
	//fmt.Println("We want the interface name that has the current IP address")
	//fmt.Println("MUST NOT be binded to 127.0.0.1 ")
	//fmt.Println("------------------------------")

	// get all the system's or local machine's network interfaces

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				//fmt.Println("[", index, "]", interf.Name, ">", addr)

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), currentIP) {
					//fmt.Println("Use name : ", interf.Name)
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	//fmt.Println("------------------------------")

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		fmt.Println(err)
	}

	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	//fmt.Println("Hardware name : ", name)
	//fmt.Println("MAC address : ", macAddress)
	//fmt.Println("OS : ", runtime.GOOS)
	//fmt.Println("CPUs : ", runtime.NumCPU())

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())

	if err != nil {
		fmt.Println("No able to parse MAC address : ", err)
		os.Exit(-1)
	}

	//fmt.Printf("Physical hardware address : %s \n", hwAddr.String())

	res, err := client.GetSystemDetails(context.Background(), &authpb.GetSystemDetailsRequest{
		HardwareName:name,
		HardwareAddress:hwAddr.String(),
		IpAddress:string(currentIP),
		Os:string(runtime.GOOS),
		Cpu:string(runtime.NumCPU()),
	})

	if err != nil{
		log.Fatalf("Error while calling GetSystemDetails RPC..!!")
	}
	log.Println("Response from GetSystemDetails RPC : ", res)
}


