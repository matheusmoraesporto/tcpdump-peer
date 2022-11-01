package sctp

import (
	"fmt"
	"unisinos/redes-i/tgb/sniffer"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	addr := fmt.Sprintf("%s:%d", ipLocal, portLocal)
	localAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro: -> MakeSCTPAddr", err)
		return
	}

	addr = fmt.Sprintf("%s:%d", ipRemote, portRemote)
	remoteAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro: -> MakeSCTPAddr", err)
		return
	}

	initMsg := NewSCTPInitMessage()
	conn, err := sctp.DialSCTP(SCTPNetowrk, localAddr, remoteAddr, &initMsg)
	if err != nil {
		fmt.Println("Erro -> DialSCTP:", err)
		return
	}
	defer conn.Close()

	sendMessageToServer(conn)
}

func sendMessageToServer(conn *sctp.SCTPConn) {
	for _, pkt := range sniffer.Sniff() {
		_, err := conn.SendMsg([]byte(pkt), nil)
		if err != nil {
			fmt.Println("Erro: -> sendMessageToServer", err)
		}
	}
}
