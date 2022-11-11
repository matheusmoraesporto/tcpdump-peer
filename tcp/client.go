package tcp

import (
	"encoding/json"
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

	return waitPackets(connection)
}

func waitPackets(connection *net.TCPConn) (packets []string) {
	buf := make([]byte, 8000)
	n, err := connection.Read(buf)
	if err != nil {
		fmt.Printf("Client side: Erro -> %s\n", err)
		return nil
	}

	if err := json.Unmarshal(buf[:n], &packets); err != nil {
		fmt.Printf("Client side: Erro -> %s\n", err)
		return nil
	}

	return
}
