package main

import (
	"fmt"
	"io" // "log"
	"net"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type fn func(...string) (string, error)

//define the IP and the port used by the TV on the network
const serverHost = "192.168.42.63"
const serverPort = "9761" //by default, this is the port used by LG

/*
TODO:
 -server
 -port
 -loglevel
*/

func main() {
	log.SetLevel(log.DebugLevel)

	m := map[string]fn{
		"mute":   mute,
		"volume": volume,
	}

	f := m[os.Args[1]]
	if f == nil {
		fmt.Printf("Error: unable to find command %s", os.Args[1])
	}

	command, err := f(os.Args[2:]...)

	if err != nil {
		log.Fatalf(`Error: unable retrieve command for "%s"`, os.Args[1])
	}

	err = sendCommand(serverHost, serverPort, command)

	if err != nil {
		log.Error("Failed to send command: %v", err)
	}
}

//This function initiate a connection between the computer and the TV
//It takes in parameters a string composed of the TV's IP and the port used
func sendCommand(srv string, port string, command string) error {
	r := strings.NewReader(command + "\n")

	log.Debugf("connecting to TV at %s:%s", srv, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", srv, port))
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return nil
	}

	log.Debugf(`sending command "%s" to TV`, command)
	_, err = io.Copy(conn, r)

	if err != nil {
		return fmt.Errorf("Connection error: %s\n", err)
	}

	return nil
}

func mute(vals ...string) (string, error) {
	if vals[0] == "" {
		return "", fmt.Errorf("invalid mute value provided: %s", vals[0])
	}

	if vals[0] == "true" {
		return "ke 00 00", nil
	} else {
		return "ke 00 01", nil
	}
}

func volume(vals ...string) (string, error) {
	value, err := strconv.ParseInt(vals[0], 10, 64)

	if err != nil {
		return "", err
	}

	if value < 0 || value > 100 {
		return "", fmt.Errorf("invalid value: %s", value)
	}
	return fmt.Sprintf("kf 00 %x", value), nil
}
