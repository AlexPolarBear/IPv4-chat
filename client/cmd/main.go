package main

import (
	"fmt"
	"net"
	"os"

	"github.com/my/repo/Desktop/IPv4chat/IP4-chat/client/internal/writer"
)

// This method connect to UDP server using a
// entering host and port. From this method the client can
// send a message in UDP chanel.
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		fmt.Scan(&arguments)
	}
	address := arguments[1]

	ser, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		fmt.Println("Connection to UDP address failed:", err.Error())
		os.Exit(1)
	}

	writer.Writer(ser)
}
