package udp

import (
	"bufio"
	"fmt"
	"net"
)

func (_ ConnectionUDP) RunClient(ip string, port int) {
	buf := make([]byte, BufferLength)
	netConn, err := net.Dial(UDP, fmt.Sprintf("%s:%d", ip, port))
	defer netConn.Close()
	if err != nil {
		fmt.Printf("Ocorreu um error: %v", err)
		panic(err)
	}

	fmt.Fprintf(netConn, RequestSniff)
	if _, err = bufio.NewReader(netConn).Read(buf); err != nil {
		fmt.Printf("Ocorreu um error: %v\n", err)
		panic(err)
	}
}
