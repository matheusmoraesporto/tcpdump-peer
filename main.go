package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"unisinos/redes-i/tgb/address"
)

func main() {
	protocolFlag := flag.String("protocol", "", "")
	flag.Parse()

	protocol := NewProtocol(*protocolFlag)
	if protocol == nil {
		fmt.Printf("Protocolo não identificado: %s\n", *protocolFlag)
		return
	}

	local, remotes, err := address.GetConnections("./address/addresses.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	go protocol.RunServer(local.Ip, local.ServerPort, remotes)
	wg.Add(1)

	select {
	case <-time.After(time.Second * 10):
	}

	for _, c := range remotes {
		go protocol.RunClient(local.Ip, c.Ip, local.ClientPort, c.ServerPort)
		wg.Add(1)
	}

	wg.Wait()
}
