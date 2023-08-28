package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Using a maximum size of buffer-massage.
const BufferSize = 1000

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

	cli, err := net.DialUDP("udp4", nil, ser)
	if err != nil {
		fmt.Println("Listen failed:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("The UDP server is %s\n", cli.RemoteAddr().String())
	defer cli.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err := cli.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println("Write data failed:", err.Error())
			os.Exit(1)
		}

		// buffer to get data
		buffer := make([]byte, BufferSize)
		n, _, err := cli.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read data faild:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
