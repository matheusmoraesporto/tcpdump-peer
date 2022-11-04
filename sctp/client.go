package sctp

import (
	"fmt"

	sctp "github.com/thebagchi/sctp-go"
)

const lenBuffer = 1500

func (_ ConnectionSCTP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) []string {
	addr := fmt.Sprintf("%s:%d", ipLocal, portRemote)
	localAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Printf("Erro: -> %s\n", err)
		return nil
	}

	addr = fmt.Sprintf("%s:%d", ipRemote, portRemote)
	remoteAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Printf("Erro: -> %s\n", err)
		return nil
	}

	initMsg := NewSCTPInitMessage()
	conn, err := sctp.DialSCTP(SCTPNetowrk, localAddr, remoteAddr, &initMsg)
	if err != nil {
		fmt.Printf("Erro: -> %s\n", err)
		return nil
	}

	defer conn.Close()
	return waitPackets(conn)
}

func waitPackets(conn *sctp.SCTPConn) (packets []string) {
	buffer := make([]byte, lenBuffer)
	flag := 0
	for {
		info := &sctp.SCTPSndRcvInfo{}
		n, err := conn.RecvMsg(buffer, info, &flag)
		if err != nil {
			fmt.Println(err)
		} else if n == 0 {
			break
		} else {
			packets = append(packets, string(buffer[:n]))
		}
	}
	return
}
