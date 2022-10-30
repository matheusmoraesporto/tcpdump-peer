package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"unisinos/redes-i/tgb/connection"
)

func main() {
	protocolFlag := flag.String("protocol", "", "")
	flag.Parse()

	protocol := NewProtocol(*protocolFlag)
	if protocol == nil {
		fmt.Printf("Protocolo não identificado: %s\n", *protocolFlag)
		return
	}

	server, clients, err := connection.GetConnections("./connection/addresses.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	go protocol.RunServer(server.Ip, server.ServerPort, clients)
	wg.Add(1)

	select {
	case <-time.After(time.Second * 5):
	}

	for _, c := range clients {
		fmt.Printf("Ip=%s\nPort=%d", c.Ip, c.ClientPort)
		go protocol.RunClient(c.Ip, c.ClientPort)
		wg.Add(1)
	}

	wg.Wait()
}
