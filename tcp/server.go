package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
	"unisinos/redes-i/tgb/connection"
)

func (_ ConnectionTCP) RunServer(ip string, port int, responseConnections []connection.Connection) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
