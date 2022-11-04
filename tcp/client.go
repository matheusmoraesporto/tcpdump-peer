package tcp

import (
	"fmt"
	"net"
)

func (_ ConnectionTCP) RunClient(ipLocal, ipRemote string, port int) []string {
	localAddr := HandleTCPAddress(ipLocal, port)
	remoteAddr := HandleTCPAddress(ipRemote, port)

	connection, err := net.DialTCP(TCPProtocol, localAddr, remoteAddr)
	if err != nil {
		fmt.Printf("Client side: Errro -> %s\n", err)
		return nil
	}

	defer func() {
		if err := connection.Close(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Client side: conexão encerrada com o endereço %s\n", remoteAddr.IP)
		}
	}()

	// escrevendo a mensagem na conexão (socket)
	if _, err := fmt.Fprintf(connection, fmt.Sprintf("teste %s\n", localAddr.String())); err != nil {
		fmt.Printf("Client side: Erro -> %s\n", err)
		return nil
	}

	packets := make([]string, 10)

	// TODO: Adicionar um listener para o servidor, que receberá os pacotes

	return packets
}
