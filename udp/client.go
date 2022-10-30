package udp

import (
	"bufio"
	"fmt"
	"net"
)

func (_ ConnectionUDP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	buf := make([]byte, BufferLength)
	netConn, err := net.Dial(UDP, fmt.Sprintf("%s:%d", ipRemote, portRemote))
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
