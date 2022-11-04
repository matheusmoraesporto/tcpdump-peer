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
		if err != nil {
			fmt.Printf("Server side: Erro -> %s\n", err)
			return
		}

		fmt.Printf("Conexão estabelecida com %s\n", connection.RemoteAddr().String())
		go sniffAndSend(connection)
	}
}

func sniffAndSend(connection *net.TCPConn) {
	for i, pkt := range sniffer.Sniff() {
		// escrevendo a mensagem na conexão (socket)
		if _, err := connection.Write([]byte(pkt)); err != nil {
			fmt.Printf("Server side: Erro -> %s\n", err)
			return
		} else {
			fmt.Printf("Enviado pacote %d\n", i)
		}
	}
}
