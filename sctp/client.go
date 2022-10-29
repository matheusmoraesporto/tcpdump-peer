package sctp

import (
	"fmt"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunClient(senderIp string, senderPort int) {
	addr := fmt.Sprintf("%s:%d", "127.0.0.3", 1234) // TODO
	clientAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	addr = fmt.Sprintf("%s:%d", senderIp, senderPort)
	serverAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	initMsg := NewSCTPInitMessage()
	conn, err := sctp.DialSCTP(SCTPNetowrk, clientAddr, serverAddr, &initMsg)
	if err != nil {
		fmt.Printf("clientAddr = %v\nserverAddr = %v\n", clientAddr, serverAddr)
		// TA DANDO ERRO AQUI
		fmt.Println("deuErro:", err)
		return
	}
	defer conn.Close()

	sendMessageToServer(conn)
}

func sendMessageToServer(conn *sctp.SCTPConn) {
	length, err := conn.SendMsg([]byte("HELLO WORLD"), nil)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Sent %d bytes\n", length)
	}
}
