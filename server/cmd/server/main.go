package main

import (
	"fmt"
	"net"
	"os"

	"github.com/my/repo/Desktop/IPv4chat/IP4-chat/server/internal/reader"
)

// This method create a UDP connection with IP4 adrress
// using a entering port.
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number.")
		fmt.Scan(&arguments)
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println("Create connection failed:", err.Error())
		os.Exit(1)
	}

	reader.Reader(s)
}
