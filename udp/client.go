package udp

import (
	"bufio"
	"fmt"
	"net"
)

func (_ ConnectionUDP) RunClient(ipLocal, ipRemote string, port int) []string {
	buf := make([]byte, BufferLength)
	netConn, err := net.Dial(UDP, fmt.Sprintf("%s:%d", ipRemote, port))
	defer netConn.Close()
	if err != nil {
		fmt.Printf("Ocorreu um error: %v", err)
		return nil
	}

	fmt.Fprintf(netConn, RequestSniff)
	if _, err = bufio.NewReader(netConn).Read(buf); err != nil {
		fmt.Printf("Ocorreu um error: %v\n", err)
		return nil
	}

	return make([]string, 0)
}
