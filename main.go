package main

import (
	"fmt"
	"sync"
	"unisinos/redes-i/tgb/client"
	"unisinos/redes-i/tgb/connections"
	"unisinos/redes-i/tgb/server"
)

func main() {
	connList, err := connections.GetConnections("./connections/addresses.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	localIp, err := connections.GetLocalIp()
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for _, c := range connList {
		if c.Ip == localIp {
			go server.RunServer(c.Ip, c.ServerPort)
		} else {
			go client.RunClient(c.Ip, c.ClientPort)
		}
		wg.Add(1)
	}
	wg.Wait()
}
