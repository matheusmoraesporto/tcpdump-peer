package tcp

import (
	"fmt"
	"net"
	"unisinos/redes-i/tgb/address"
	"unisinos/redes-i/tgb/sniffer"
)

func (_ ConnectionTCP) RunServer(ip string, port int, responseAddresses []address.Address) {
	addr := HandleTCPAddress(ip, port)
	listener, err := net.ListenTCP(TCPProtocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Printf("Servidor executando no endereço %s\n", addr.String())

	for {
		connection, err := listener.AcceptTCP()
		clientaddr := connection.RemoteAddr().String()
		fmt.Printf("Conexão estabelecida com %s\n", connection.RemoteAddr().String())
		if err != nil {
			fmt.Println(err)
			return
		}

		sniffAndSend(connection)

		if err := connection.Close(); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Printf("Server side: conexão encerrada com o client %s\n", clientaddr)
		}
	}
}

func sniffAndSend(connection *net.TCPConn) {
	for _, pkt := range sniffer.Sniff() {
		// escrevendo a mensagem na conexão (socket)
		if _, err := fmt.Fprintf(connection, pkt); err != nil {
			fmt.Printf("Client side: Erro -> %s\n", err)
			return
		}
	}
}
