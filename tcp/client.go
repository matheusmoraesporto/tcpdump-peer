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

	// for {
	// escrevendo a mensagem na conexão (socket)
	fmt.Fprintf(connection, fmt.Sprintf("teste %s\n", localAddr.String()))

	// ouvindo a resposta do servidor (eco)
	bufio.NewReader(connection).ReadString('\n')
	// }
}
