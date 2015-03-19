// usage: byron e805bd8086eebade820ac8368ec2fbf4
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var client http.Client

var mode = flag.String("mode", "", "command to run")
var token = flag.String("token", "", "API token")
var user = flag.String("user", "", "user to run as")

func init() {
	flag.StringVar(mode, "m", "", "command to run")
	flag.StringVar(token, "t", "", "API token")
	flag.StringVar(user, "u", "", "user to run as")
	flag.Parse()
}

func usage() {
	fmt.Println("usage: byron --t <token> --m <command>")
	os.Exit(1)
}

func main() {
	client := newClient()
	switch *mode {
	case "install":
		installAgent(client, *token, *user)
	case "remove":
		removeAgent(client, *token)
	default:
		usage()
	}
}
