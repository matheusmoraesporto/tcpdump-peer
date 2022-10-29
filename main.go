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
			go protocol.RunServer(c.Ip, c.ServerPort, connList)
			wg.Add(1)
			break
		}
	}

	time.Sleep(time.Second * 10)

	for _, c := range connList {
		if c.Ip != localAddr.Ip {
			go protocol.RunClient(c.Ip, c.ServerPort)
			wg.Add(1)
		}
	}

	wg.Wait()
}
