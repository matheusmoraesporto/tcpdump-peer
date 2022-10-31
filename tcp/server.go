package tcp

import (
	"bufio"
	"fmt"
	"net"
	"unisinos/redes-i/tgb/address"
)

func (_ ConnectionTCP) RunServer(ip string, port int, responseAddresses []address.Address) {
	addr := HandleTCPAddress(ip, port)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Printf("Servidor executando no endereço %s\n", addr.String())

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("MENSAGEM RECEBIDA: ", string(netData))
		connection.Close()
	}
}
