package sctp

import (
	"errors"
	"fmt"
	"time"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) []string {
	addr := fmt.Sprintf("%s:%d", ipLocal, portLocal)
	localAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro: -> MakeSCTPAddr", err)
		return nil
	}

	addr = fmt.Sprintf("%s:%d", ipRemote, portRemote)
	remoteAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro: -> MakeSCTPAddr", err)
		return nil
	}

	initMsg := NewSCTPInitMessage()
	conn, err := retryConnection(localAddr, remoteAddr, &initMsg)
	if err != nil {
		return nil
	}
	defer conn.Close()
	return waitPackets(conn)
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

func waitPackets(conn *sctp.SCTPConn) (packets []string) {
	data := make([]byte, 8192)
	flag := 0
	for {
		info := &sctp.SCTPSndRcvInfo{}
		n, err := conn.RecvMsg(data, info, &flag)
		fmt.Println(n)
		if err != nil {
			fmt.Println(err)
		} else if info != nil && info.Flags == sctp.SCTP_SHUTDOWN_SENT {
			break
		} else {
			packets = append(packets, string(data[:n]))
		}
	}
	return
}
