package main

import (
	"fmt"
	"sync"
	"time"
	"unisinos/redes-i/tgb/address"
)

func main() {
	// protocolFlag := flag.String("protocol", "", "")
	// flag.Parse()

	protocol := NewProtocol("sctp")
	// if protocol == nil {
	// 	fmt.Printf("Protocolo não identificado: %s\n", *protocolFlag)
	// 	return
	// }

	local, remotes, err := address.GetConnections("./address/addresses.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	go protocol.RunServer(local.Ip, local.ServerPort, remotes)

	select {
	case <-time.After(time.Second * 10):
	}

	mut := new(sync.Mutex)
	pktsByAddress := make(map[string][]string)
	for _, r := range remotes {
		go func(r address.Address) {
			pkts := protocol.RunClient(local.Ip, r.Ip, local.ClientPort, r.ServerPort)
			fmt.Printf("Solicitando pacotes para o servidor %s:%d\n", r.Ip, r.ClientPort)

			mut.Lock()
			pktsByAddress[r.Ip] = pkts
			mut.Unlock()
		}(r)

		wg.Add(1)
	}

	wg.Wait()

	printReceivedData(pktsByAddress)

	// Adiciona para deixar o terminal bloqueante, pois o server pode ainda receber alguma conexão
	wg.Add(1)
	wg.Wait()
}

func printReceivedData(packetsByAddress map[string][]string) {
	for addr, packets := range packetsByAddress {
		fmt.Printf("Pacotes recebidos pelo endereço %s", addr)
		fmt.Println("--------------------------------------------")
		for _, pkt := range packets {
			fmt.Println(pkt)
			fmt.Println("--------------------------------------------")
		}
	}
}
