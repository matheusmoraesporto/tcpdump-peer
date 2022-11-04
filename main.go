package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"unisinos/redes-i/tgb/address"
)

const (
	addressesFilePath      = "./address/addresses.json"
	protocolFlag           = "protocol"
	timeWaitToCreateServer = 10
)

func main() {
	flagValue := flag.String(protocolFlag, "", "")
	flag.Parse()

	protocol := NewProtocol(*flagValue)
	if protocol == nil {
		fmt.Printf("Protocolo não identificado: %s\n", *flagValue)
		return
	}

	local, remotes, err := address.GetAddresses(addressesFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go protocol.RunServer(local.Ip, local.Port, remotes)

	select {
	case <-time.After(time.Second * timeWaitToCreateServer):
	}

	pktsByAddress := requestSniff(protocol, local, remotes)
	printReceivedData(pktsByAddress)
	wg.Wait()
}

func requestSniff(protocol Protocol, localAddr address.Address, remotes []address.Address) map[string][]string {
	pktsByAddress := make(map[string][]string)
	wg := new(sync.WaitGroup)
	mut := new(sync.Mutex)

	for _, r := range remotes {
		wg.Add(1)
		go func(r address.Address) {
			defer wg.Done()
			pkts := protocol.RunClient(localAddr.Ip, r.Ip, r.Port)
			fmt.Printf("Solicitando pacotes para o servidor %s:%d\n", r.Ip, r.Port)

			mut.Lock()
			pktsByAddress[r.Ip] = pkts
			mut.Unlock()
		}(r)
	}

	wg.Wait()
	return pktsByAddress
}

func printReceivedData(packetsByAddress map[string][]string) {
	for addr, packets := range packetsByAddress {
		fmt.Printf("\n\nPacotes recebidos pelo endereço %s\n", addr)
		fmt.Println("--------------------------------------------")
		for _, pkt := range packets {
			fmt.Println(pkt)
			fmt.Println("--------------------------------------------")
		}
	}
}
