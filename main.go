package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
	//"os"
)

//define the IP and the port used by the TV on the network
const serverHost = "192.168.42.63"
const serverPort = "9761" //by default, this is the port used by LG

//Manage command line arguments
var (
	mute = flag.String("mute", "", "Boolean value : true ==  muted")
)

func main() {
	flag.Parse()

	fmt.Println("muted:", *mute)

	if *mute != "" {
		if *mute == "true" {
			startClient(fmt.Sprintf("%s:%s", serverHost, serverPort), "ke 00 00\n")
		} else {
			startClient(fmt.Sprintf("%s:%s", serverHost, serverPort), "ke 00 01\n")
		}
	}

}

//This function initiate a connection between the computer and the TV
//It takes in parameters a string composed of the TV's IP and the port used
func startClient(addr string, command string) {
	fmt.Println("Trying to connect to the TV...")

	r := strings.NewReader(command)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return
	}
	_, err = io.Copy(conn, r)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}

	fmt.Println("End")
}
