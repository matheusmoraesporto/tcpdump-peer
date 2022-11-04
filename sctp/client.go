package sctp

import (
	"fmt"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) []string {
	addr := fmt.Sprintf("%s:%d", ipLocal, portRemote)
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
	conn, err := sctp.DialSCTP(SCTPNetowrk, localAddr, remoteAddr, &initMsg)
	if err != nil {
		fmt.Println("Erro -> DialSCTP:", err)
		return nil
	}

	defer conn.Close()
	return waitPackets(conn)
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
