package tcp

import "net"

type ConnectionTCP struct{}

func HandleTCPAddress(ip string, port int) *net.TCPAddr {
	if port == 0 {
		return &net.TCPAddr{
			IP: net.ParseIP(ip),
		}
	}

	return &net.TCPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
}
