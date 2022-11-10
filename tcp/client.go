package tcp

import (
	"encoding/json"
	"fmt"
	"io"
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

	return waitPackets(connection)
}

func waitPackets(connection *net.TCPConn) (packets []string) {
	buf := make([]byte, 15000)
	fmt.Println("Client: chamou o ReadFull")
	_, err := io.ReadFull(connection, buf)
	fmt.Println("Client: passou do ReadFull")
	if err != nil {
		fmt.Printf("Client side: Erro -> %s\n", err)
		return nil
	}

	fmt.Println("Client: chamou o Unmarshal")
	if err := json.Unmarshal(buf, &packets); err != nil {
		fmt.Printf("Client side: Erro -> %s\n", err)
		return nil
	}

	fmt.Println("Client: passou do Unmarshal e irá retornar os pacotes")
	return packets
}
