package tcp

import (
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
	buf := make([]byte, 1500)
	for {
		_, err := io.ReadFull(connection, buf)
		if err != nil {
			fmt.Printf("Client side: Erro -> %s\n", err)
			return nil
		}
		packets = append(packets, string(buf))

		if len(packets) == 10 {
			break
		} else {
			fmt.Printf("packets len = %d\n", len(packets))
		}
	}
	return packets
}
