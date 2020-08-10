package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var serverHost, serverPort string

func main() {
	app := cli.NewApp()
	app.Name = "LG TV command-line interface"
	app.Usage = "Let's you drive your LG TV from your terminal!"
	app.EnableBashCompletion = true

	// set the viper options to read config.yaml
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// check if config.yaml exists
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Impossible to read conf file: %v", err)
	}

	serverHost = inConfigFile("ip").(string)
	serverPort = inConfigFile("port").(string)

	app.Commands = []cli.Command{
		{
			Name:  "volume",
			Usage: "Adjusts the TV volume. The value must be in 0-100",
			Action: func(c *cli.Context) error {
				volVal := c.Args().First()
				volValInt, _ := strconv.Atoi(volVal)
				err := sendCommand(serverHost, serverPort, fmt.Sprintf("kf 00 %.2x", volValInt))
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "mute",
			Usage: "Turns on/off the TV sound. The argument must be true to mute, anything else to unmute",
			Action: func(c *cli.Context) error {
				if strings.ToLower(c.Args().First()) == "true" || c.Args().First() == "1" {
					// mute the sound
					err := sendCommand(serverHost, serverPort, "ke 00 00")
					if err != nil {
						return err
					}
					return nil
				}

				// unmute the sound
				err := sendCommand(serverHost, serverPort, "ke 00 01")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "poweroff",
			Usage: "Turns off the TV. Needs no arguments",
			Action: func(c *cli.Context) error {
				err := sendCommand(serverHost, serverPort, "ka 00 00")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "brightness",
			Usage: "Adjusts the brightness of the screen. The value must be in 0-100",
			Action: func(c *cli.Context) error {
				brightVal := c.Args().First()
				brightValInt, _ := strconv.Atoi(brightVal)
				err := sendCommand(serverHost, serverPort, fmt.Sprintf("kh 00 %.2x", brightValInt))
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	// Sort commands list in help panel by name
	sort.Sort(cli.CommandsByName(app.Commands))

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func inConfigFile(param string) interface{} {
	if viper.Get(param) == nil {
		return fmt.Sprintf("%s not found in config file", param)
	}

	return viper.Get(param)
}

//This function sends a command string to the IP address given at the given port
func sendCommand(srv string, port string, command string) error {

	// add the command to a new reader for io and append an necessited carriage return
	r := strings.NewReader(command + "\n")

	logrus.Debugf("connecting to TV at %s:%s", srv, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", srv, port))
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return nil
	}

	logrus.Debugf("sending command %s to TV", command)

	// copy the reader in the io, so this "send" it to the server (copy-paste to the i buffer)
	_, err = io.Copy(conn, r)

	if err != nil {
		return fmt.Errorf("Connection error: %d", err)
	}

	return nil
}
