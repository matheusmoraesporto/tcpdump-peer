package sniffer

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	device         = "eth0"
	defaultSnapLen = 262144
	tcpProtocol    = "TCP" // TODO: talvez essas duas constantes sejam removidas, uma vez que os logs poderão mudar
	udpProtocol    = "UDP"
)

// TODO: também poderá ser deletado
type SniffedPacket struct {
	Checksum       uint16 `json:"checksum"`
	DestinyAddress string `json:"destinyAddress"`
	IpType         string `json:"ipType"`
	Length         int    `json:"length"`
	OriginAddress  string `json:"originAddress"`
	Protocol       string `json:"protocol"`
}

func Sniff() (spkt SniffedPacket) {
	handle, err := pcap.OpenLive(device, defaultSnapLen, false, time.Duration(time.Minute))
	if err != nil {
		fmt.Printf("Erro ao iniciar processo de captura de pacotes: %v", err)
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		spkt = handleData(pkt)
		break
	}

	return
}

func handleData(pkt gopacket.Packet) SniffedPacket {
	networkFlow := pkt.NetworkLayer().NetworkFlow()

	sp := SniffedPacket{
		DestinyAddress: networkFlow.Dst().String(),
		IpType:         networkFlow.EndpointType().String(),
		OriginAddress:  networkFlow.Src().String(),
		Length:         len(pkt.Data()),
	}

	handleUDP(&sp, pkt)
	handleTCP(&sp, pkt)

	return sp
}

func handleTCP(sp *SniffedPacket, pkt gopacket.Packet) {
	if tcpLayer := pkt.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)

		sp.Checksum = tcp.Checksum
		sp.Protocol = tcpProtocol
	}
}

func handleUDP(sp *SniffedPacket, pkt gopacket.Packet) {
	if udpLayer := pkt.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)

		sp.Checksum = udp.Checksum
		sp.Protocol = udpProtocol
	}
}
