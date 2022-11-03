package tcp

import "net"

const TCPProtocol = "tcp"

type ConnectionTCP struct{}

func HandleTCPAddress(ip string, port int) *net.TCPAddr {
	return &net.TCPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
}
