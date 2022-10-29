package udp

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"unisinos/redes-i/tga/common"
	"unisinos/redes-i/tgb/sniffer"
)

const lenSniffPackets = 10

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
		n, remoteaddr, err := server.ReadFromUDP(buf) // bloqueante
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
			var pkt sniffer.SniffedPacket
			json.Unmarshal(buf[:n], &pkt)
			printSniffedPackets(remoteaddr, pkt)
		}
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	for i := 0; i < lenSniffPackets; i++ {
		pkts := sniffer.Sniff()
		pktJson, err := json.Marshal(pkts)
		if err != nil {
			fmt.Printf("Erro ao tratar os dados sniffados: %v", err)
			return
		}

		_, err = conn.WriteToUDP(pktJson, addr)
		if err != nil {
			fmt.Printf("Erro ao enviar a resposta: %v", err)
			return
		}

		fmt.Println("Pacote enviado.")
	}
}

func printSniffedPackets(addr *net.UDPAddr, pkt sniffer.SniffedPacket) {
	fmt.Println("=============================================================")
	fmt.Printf("Pacote sniffado e recebido pelo endereço: %s\n\n", addr.IP)
	fmt.Printf("Endereço que enviou o pacote: %s\n", pkt.OriginAddress)
	fmt.Printf("Endereço que recebeu o pacote: %s\n", pkt.DestinyAddress)
	fmt.Printf("Protocolo da mensagem: %s\n", pkt.Protocol)
	fmt.Printf("Tamanho do pacote: %dbytes\n", pkt.Length)
	fmt.Printf("Tipo de IP: %s\n", pkt.IpType)
	fmt.Printf("Checksum do cabeçalho de transporte: %d\n", pkt.Checksum)
	fmt.Println("=============================================================")
}
