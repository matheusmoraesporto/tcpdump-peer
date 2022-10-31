package sctp

import (
	"fmt"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	addr := fmt.Sprintf("%s:%d", ipLocal, portLocal)
	clientAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	addr = fmt.Sprintf("%s:%d", ipRemote, portRemote)
	serverAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	initMsg := NewSCTPInitMessage()
	conn, err := sctp.DialSCTP(SCTPNetowrk, clientAddr, serverAddr, &initMsg)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer conn.Close()

	sendMessageToServer(conn)
}

func sendMessageToServer(conn *sctp.SCTPConn) {
	msg := fmt.Sprintf("Oi eu sou %s", conn.LocalAddr().String())
	_, err := conn.SendMsg([]byte(msg), nil)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
