package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// This method create a UDP connection with IP4 adrress
// using a entering port.
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(1)
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read data failed", err.Error())
			os.Exit(1)
		}
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			os.Exit(1)
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
