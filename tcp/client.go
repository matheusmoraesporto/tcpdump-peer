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
	fmt.Printf("connection | %v", connection)

	if _, err := bufio.NewReader(connection).ReadString('\n'); err != nil {
		fmt.Printf("Ocorreu um error: %v\n", err)
		panic(err)
	}
	// if strings.TrimSpace(string(text)) == "STOP" {
	// 	fmt.Println("TCP client exiting...")
	// 	return
	// }
	// }
}
