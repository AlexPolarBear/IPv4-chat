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
		fmt.Println("Listen failed: ", err.Error())
		os.Exit(1)
	}

	nickname := setname.SetName()

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	fmt.Println("Name: " + nickname)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		msg := "User: " + nickname + "\n" + "Msg: " + text + "\n\n"
		data := []byte(msg)
		_, err := c.Write(data)
		if err != nil {
			fmt.Println("Write data failed: ", err.Error())
			os.Exit(1)
		}

		if strings.Contains(strings.TrimSpace(string(data)), "STOP") {
			fmt.Println("Exiting UDP client!")
			return
		}
	}
}
