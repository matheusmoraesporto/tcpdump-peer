package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func (_ ConnectionTCP) RunClient(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	c, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
