package writer

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/my/repo/Desktop/IPv4chat/IP4-chat/client/internal/setname"
)

// Using a maximum size of buffer-massage.
const BufferSize = 1000

// Connect to UDP and send massages
func Writer(s *net.UDPAddr) {
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println("Listen failed:", err.Error())
		os.Exit(1)
	}

	nickname := setname.SetName()

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Name: " + nickname)
		fmt.Print(">>")
		text, _ := reader.ReadString('\n')
		data := []byte("Name: " + nickname + "\n" + text + "\n")
		_, err := c.Write(data)
		if err != nil {
			fmt.Println("Write data failed:", err.Error())
			os.Exit(1)
		}
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		// buffer to get data
		buffer := make([]byte, BufferSize)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read data faild:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
