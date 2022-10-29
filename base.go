package main

import (
	"unisinos/redes-i/tgb/sctp"
	"unisinos/redes-i/tgb/tcp"
	"unisinos/redes-i/tgb/udp"
)

type Protocol interface {
	RunClient(ip string, port int)
	RunServer(ip string, port int, portResponse int)
}

func NewTCP() (p Protocol) {
	return tcp.ConnectionTCP{}
}

func NewUDP() (p Protocol) {
	return udp.ConnectionUDP{}
}

func NewSCTP() (p Protocol) {
	return sctp.ConnectionSCTP{}
}
