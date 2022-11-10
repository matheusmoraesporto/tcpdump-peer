package tcp

import (
	"encoding/json"
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
	// defer connection.Close()

	pkts := sniffer.Sniff()
	fmt.Printf("LEN PKTS = %d", len(pkts))
	buffer, err := json.Marshal(pkts)
	if err != nil {
		fmt.Printf("Server side: Erro -> %s\n", err)
		return
	}

	if _, err := connection.Write([]byte(buffer)); err != nil {
		fmt.Printf("Server side: Erro -> %s\n", err)
		return
	}

	fmt.Printf("Todos os pacotes foram enviados para %s\n", connection.RemoteAddr())
}
