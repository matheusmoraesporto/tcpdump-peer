package tcp

import (
	"errors"
	"fmt"
	"net"
	"time"
)

const msgAddressAlreadyInUse = "bind: address already in use"

func (_ ConnectionTCP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	localAddr := HandleTCPAddress(ipLocal, portLocal)
	remoteAddr := HandleTCPAddress(ipRemote, portRemote)

	connection, err := retryConnection(localAddr, remoteAddr)
	if err != nil {
		return
	}

	defer connection.Close()

	// for {
	// escrevendo a mensagem na conexão (socket)
	fmt.Fprintf(connection, fmt.Sprintf("teste %s\n", localAddr.String()))

	// ouvindo a resposta do servidor (eco)
	// bufio.NewReader(connection) //.ReadString('\n')
	// }
}

func retryConnection(localAddr *net.TCPAddr, remoteAddr *net.TCPAddr) (*net.TCPConn, error) {
	retryfunc := func() (*net.TCPConn, error) {
		connection, err := net.DialTCP("tcp", localAddr, remoteAddr)
		if err != nil {
			fmt.Println("deu erro aqui")
			fmt.Println(err)
			return nil, err
		}

		return connection, err
	}

	retryPeriod := time.Second * 10
	timeout := time.After(time.Second * 45)
	for {
		select {
		case <-timeout:
			connection, err := retryfunc()

			if err == nil {
				return connection, err
			}
			return nil, errors.New("Máxima tentativa de conxões atingidas")
		case <-time.After(retryPeriod):
		}
	}
}
