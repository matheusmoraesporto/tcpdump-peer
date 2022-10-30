package tcp

import (
	"fmt"
	"net"
)

func (_ ConnectionTCP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	localAddr := HandleTCPAddress(ipLocal, portLocal)
	remoteAddr := HandleTCPAddress(ipRemote, portRemote)

	var connection *net.TCPConn
	var err error
	for {
		connection, err = net.DialTCP("tcp", localAddr, remoteAddr)
		if err == nil {
			break
		}
	}

	defer connection.Close()

	// for {
	// escrevendo a mensagem na conexão (socket)
	fmt.Fprintf(connection, fmt.Sprintf("teste %s\n", localAddr.String()))

	// ouvindo a resposta do servidor (eco)
	// bufio.NewReader(connection) //.ReadString('\n')
	// }
}
