package udp

import (
	"fmt"
	"net"
	"strings"
	"unisinos/redes-i/tgb/address"
	"unisinos/redes-i/tgb/sniffer"
)

func (_ ConnectionUDP) RunServer(ip string, port int, responseAddresses []address.Address) {
	address := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}

	server, err := net.ListenUDP(UDP, &address)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao iniciar o server para o endereço %v:%d: %v\n", &address.IP, address.Port, err)
		return
	}

	buf := make([]byte, BufferLength)
	fmt.Printf("Servidor rodando no endereço %s:%d\n", ip, port)
	for {
		_, remoteaddr, err := server.ReadFromUDP(buf) // bloqueante
		if err != nil {
			fmt.Printf("Ocorreu um erro: %v", err)
			continue
		}

		// se o buffer for GET significa que há um servidor pedindo para ser sniffado
		// então podemos iniciar o sniffer para enviar a response
		if strings.Contains(string(buf), RequestSniff) {
			fmt.Printf("O endereço %s:%d solicitou sniffer de pacotes\n", remoteaddr.IP.String(), remoteaddr.Port)
			for _, a := range responseAddresses {
				if a.Ip == remoteaddr.IP.String() {
					remoteaddr.Port = a.ServerPort
				}
			}
			sendResponse(server, remoteaddr)
		} else { // caso contrário, estamos recebendo os dados sniffados de outro servidor
			printSniffedPackets(remoteaddr, string(buf))
		}
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	// go func() {
	for _, pkt := range sniffer.Sniff() {
		if _, err := conn.WriteToUDP([]byte(pkt), addr); err != nil {
			fmt.Printf("Erro ao enviar a resposta: %v", err)
			return
		}

		fmt.Printf("Pacote enviado para o endereço %s:%d.\n", addr.IP, addr.Port)
	}
	// }()
}

func printSniffedPackets(addr *net.UDPAddr, pkt string) {
	fmt.Println("=============================================================")
	fmt.Printf("Pacote sniffado e recebido pelo endereço: %s\n\n%s\n", addr.IP, pkt)
	fmt.Println("=============================================================")
}
