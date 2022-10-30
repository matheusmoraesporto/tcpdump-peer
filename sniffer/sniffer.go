package sniffer

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	device          = "lo"
	defaultSnapLen  = 262144
	lenSniffPackets = 10
)

func Sniff() (pkts []string) {
	handle, err := pcap.OpenLive(device, defaultSnapLen, false, time.Duration(time.Minute))
	if err != nil {
		fmt.Printf("Erro ao iniciar processo de captura de pacotes: %v", err)
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		pkts = append(pkts, pkt.String())

		if len(pkts) == lenSniffPackets {
			break
		}
	}

	return
}
