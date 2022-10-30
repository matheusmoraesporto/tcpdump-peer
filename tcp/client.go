package tcp

import (
	"bufio"
	"fmt"
	"net"
)

func (_ ConnectionTCP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	localAddr := HandleTCPAddress(ipLocal, 0)
	remoteAddr := HandleTCPAddress(ipRemote, portRemote)
	connection, err := net.DialTCP("tcp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Comunicando-se com o endereço %s\n", connection.RemoteAddr())
	// for {
	fmt.Fprintf(connection, fmt.Sprintf("Oi eu sou %s", localAddr))

	x, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Printf("Ocorreu um error: %v\n", err)
		panic(err)
	}
	fmt.Println(x)
	// if strings.TrimSpace(string(text)) == "STOP" {
	// 	fmt.Println("TCP client exiting...")
	// 	return
	// }
	// }
}
