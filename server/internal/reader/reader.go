package reader

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// Using a maximum size of buffer-massage.
const BufferSize = 1000

func Reader(s *net.UDPAddr) {
	c, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer c.Close()
	buffer := make([]byte, BufferSize)
	// rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().Unix()))

	for {
		n, addr, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read data failed", err.Error())
			os.Exit(1)
		}
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(BufferSize))
		fmt.Printf("data: %s\n", string(data))
		_, err = c.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
