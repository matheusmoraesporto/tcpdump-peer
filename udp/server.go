package udp

import (
	"fmt"
	"net"
	"strings"

	"unisinos/redes-i/tga/common"
	"unisinos/redes-i/tgb/sniffer"
)

func (_ ConnectionUDP) RunServer(ip string, port int, portResponse int) {
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
		if strings.Contains(string(buf), common.RequestSniff) {
			fmt.Printf("O endereço %s:%d solicitou sniffer de pacotes\n", remoteaddr.IP.String(), remoteaddr.Port)
			remoteaddr.Port = portResponse
			sendResponse(server, remoteaddr)
		} else { // caso contrário, estamos recebendo os dados sniffados de outro servidor
			printSniffedPackets(remoteaddr, string(buf))
		}
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	for _, pkt := range sniffer.Sniff() {
		if _, err := conn.WriteToUDP([]byte(pkt), addr); err != nil {
			fmt.Printf("Erro ao enviar a resposta: %v", err)
			return
		}

		fmt.Println("Pacote enviado.")
	}
}

func printSniffedPackets(addr *net.UDPAddr, pkt string) {
	fmt.Println("=============================================================")
	fmt.Printf("Pacote sniffado e recebido pelo endereço: %s\n\n%s\n", addr.IP, pkt)
	fmt.Println("=============================================================")
}
