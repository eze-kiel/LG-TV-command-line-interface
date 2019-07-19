package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

//this type represent a variadic func which returns a string and the error
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

	//This map uses the fn type introduced before
	//We create a map called m which associates a string keyword (between [])
	// with a function which is build following the fn type
	m := map[string]fn{
		"mute":   mute,
		"volume": volume,
	}

	// f represents a function associated with the first argument given to the program
	// so by entering "mute" as first arg, thanks to the m map, f represents
	// the function mute(...string) (string,error)
	f := m[os.Args[1]]
	if f == nil {
		fmt.Printf("Error: unable to find command %s", os.Args[1])
	}

	// the f function is applied with the second arg given to program. It should be a value or a state to
	// apply to the TV (e.g the volume value, the brightness...)
	command, err := f(os.Args[2:]...)

	if err != nil {
		log.Fatalf(`Error: unable retrieve command for "%s"`, os.Args[1])
	}

	// the sendCommand function is used. This one just send the command to the server (TV)
	err = sendCommand(serverHost, serverPort, command)

	if err != nil {
		log.Error("Failed to send command: %v", err)
	}
}

//This function send a command string to the IP address given at the port given
func sendCommand(srv string, port string, command string) error {

	// add the command to a new reader for io and append an necessited carriage return
	r := strings.NewReader(command + "\n")

	log.Debugf("connecting to TV at %s:%s", srv, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", srv, port))
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return nil
	}

	log.Debugf(`sending command "%s" to TV`, command)

	// copy the reader in the io, so this "send" it to the server (copy-paste to the i buffer)
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
		// mute the sound
		return "ke 00 00", nil
	} else {
		// unmute the sound
		return "ke 00 01", nil
	}
}

func volume(vals ...string) (string, error) {
	// convert the string into a int
	value, err := strconv.ParseInt(vals[0], 10, 64)

	if err != nil {
		return "", err
	}

	if value < 0 || value > 100 {
		return "", fmt.Errorf("invalid value: %s", value)
	}
	// append the variable part of the function to the fixed one
	return fmt.Sprintf("kf 00 %x", value), nil
}
