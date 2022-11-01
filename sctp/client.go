package sctp

import (
	"errors"
	"fmt"
	"time"
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
	conn, err := retryConnection(localAddr, remoteAddr, &initMsg)
	if err != nil {
		return
	}
	defer conn.Close()

	sendMessageToServer(conn)
}

func retryConnection(localAddr *sctp.SCTPAddr, remoteAddr *sctp.SCTPAddr, initMsg *sctp.SCTPInitMsg) (*sctp.SCTPConn, error) {
	retryFunc := func() (*sctp.SCTPConn, error) {
		conn, err := sctp.DialSCTP(SCTPNetowrk, localAddr, remoteAddr, initMsg)
		if err != nil {
			fmt.Println("Erro -> DialSCTP:", err)
			return nil, err
		}

		return conn, err
	}

	retryPeriod := time.Second * 10
	timeout := time.After(time.Second * 45)
	for {
		select {
		case <-timeout:
			connection, err := retryFunc()
			if err == nil {
				return connection, err
			}
			return nil, errors.New("Máxima tentativa de conxões atingidas")
		case <-time.After(retryPeriod):
			connection, err := retryFunc()

			if err == nil {
				return connection, err
			}
		}
	}
}

func sendMessageToServer(conn *sctp.SCTPConn) {
	for _, pkt := range sniffer.Sniff() {
		_, err := conn.SendMsg([]byte(pkt), nil)
		if err != nil {
			fmt.Println("Erro: -> sendMessageToServer", err)
		}
	}
}
