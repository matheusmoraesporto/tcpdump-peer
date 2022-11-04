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

	wg := new(sync.WaitGroup)
	go protocol.RunServer(local.Ip, local.ServerPort, remotes)
	wg.Add(1)

	select {
	case <-time.After(time.Second * 10):
	}

	pktsByAddress := requestSniff(protocol, local, remotes)
	printReceivedData(pktsByAddress)
	wg.Wait()
}

func requestSniff(protocol Protocol, localAddr address.Address, remotes []address.Address) (pktsByAddress map[string][]string) {
	wg := new(sync.WaitGroup)
	mut := new(sync.Mutex)

	for _, r := range remotes {
		go func(r address.Address) {
			defer wg.Done()
			pkts := protocol.RunClient(localAddr.Ip, r.Ip, localAddr.ClientPort, r.ServerPort)
			fmt.Printf("Solicitando pacotes para o servidor %s:%d\n", r.Ip, r.ClientPort)

			mut.Lock()
			pktsByAddress[r.Ip] = pkts
			mut.Unlock()
		}(r)

		wg.Add(1)
	}

	fmt.Println("Aguardando os clients finalizarem")
	wg.Wait()
	fmt.Println("Clients finalizados")
	return
}

func printReceivedData(packetsByAddress map[string][]string) {
	fmt.Println("Entrou no printReceivedData")
	for addr, packets := range packetsByAddress {
		fmt.Printf("Pacotes recebidos pelo endereço %s", addr)
		fmt.Println("--------------------------------------------")
		for _, pkt := range packets {
			fmt.Println(pkt)
			fmt.Println("--------------------------------------------")
		}
	}
}
