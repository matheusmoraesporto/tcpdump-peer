package main

import (
	"unisinos/redes-i/tgb/address"
	"unisinos/redes-i/tgb/sctp"
	"unisinos/redes-i/tgb/tcp"
	"unisinos/redes-i/tgb/udp"
)

type Protocol interface {
	RunClient(ipLocal, ipRemote string, port int) []string
	RunServer(ip string, port int, responseConnections []address.Address)
}

func NewProtocol(flag string) (p Protocol) {
	switch flag {
	case "tcp":
		return tcp.ConnectionTCP{}
	case "udp":
		return udp.ConnectionUDP{}
	case "sctp":
		return sctp.ConnectionSCTP{}
	default:
		return nil
	}
}
