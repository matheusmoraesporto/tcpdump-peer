package main

import (
	"flag"
	"fmt"
	"sync"
	"unisinos/redes-i/tgb/connection"
)

func main() {
	protocolFlag := flag.String("protocol", "", "")
	flag.Parse()

	protocol := indetifyProtocol(*protocolFlag)
	if protocol == nil {
		fmt.Printf("Protocolo não identificado: %s\n", *protocolFlag)
		return
	}

	connList, err := connection.GetConnections("./connection/addresses.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	localAddr, err := connection.GetLocalAddress(connList)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for _, c := range connList {
		if c.Ip == localAddr.Ip {
			go protocol.RunServer(c.Ip, c.ServerPort, c.ClientPort)
		} else {
			go protocol.RunClient(c.Ip, c.ClientPort)
		}
		wg.Add(1)
	}
	wg.Wait()
}

func indetifyProtocol(flag string) (p Protocol) {
	switch flag {
	case "tcp":
		return NewTCP()
	case "udp":
		return NewUDP()
	case "sctp":
		return NewTCP()
	default:
		return nil
	}
}
