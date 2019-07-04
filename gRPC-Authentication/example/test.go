package main

import (
	"auth_server/authpb"
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/tj/go-spin"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}


func main(){
	runtimeOS := runtime.GOOS
	fmt.Println(runtimeOS)
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	fmt.Println(vmStat)
	path := "/"
	diskStat, err := disk.Usage(path)
	dealwithErr(err)
	fmt.Println(diskStat)
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	fmt.Println(cpuStat)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)
	fmt.Println(percentage)

	hostStat, err := host.Info()
	dealwithErr(err)
	fmt.Println(hostStat)

	interfStat, err := net.Interfaces()
	dealwithErr(err)
	fmt.Println(interfStat)

	processList, err := ps.Processes()

	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		fmt.Printf("%d\t%s\n",process.Pid(),process.Executable())

		// do os.* stuff on the pid
	}

}

func consoleInterface(){

	option := grpc.WithInsecure()

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:8080", option)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	//s := spin.New()
	//show(s, spin.Default)
	client := authpb.NewAuthServiceClient(conn)
	res := &authpb.CheckUsernameResponse{}
	prompt := promptui.Prompt{
		Label: "Username",
		Validate: func(input string) error {
			res, err =client.CheckUsername(context.Background(), &authpb.CheckUsernameRequest{Username:input})
			if err != nil {
				return errors.New(err.Error())
			}
			return nil
		},

	}

	_, err = prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d := color.New(color.FgCyan, color.Bold)
	_, _ = d.Println("Hi " + res.Username + "..!!")


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
	fmt.Println(loginResponse)

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
	excute(res.Username, password, "localhost", "3650")
}

func excute(username, password, network, expiryDays string)  {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine..!!")
		return
	}

	out, err := exec.Command( "./temp_test.sh", username, password, network, expiryDays).Output()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
}

func show(s *spin.Spinner, frames string) {
	s.Set(frames)
	for i := 0; i < 30; i++ {
		fmt.Printf("\r  \033[36mverifying\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n")
}