package client

import (
	"fmt"
	"unisinos/redes-i/tgb/common"

	sctp "github.com/thebagchi/sctp-go"
)

func main() {
	run("127.0.0.1", 54321)
}

func RunClient(ip string, port int) {
	run(ip, port)
}

func run(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	clientAddr, err := sctp.MakeSCTPAddr(common.SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	serverAddr, err := sctp.MakeSCTPAddr(common.SCTPNetowrk, "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	initMsg := common.NewSCTPInitMessage()
	conn, err := sctp.DialSCTP(common.SCTPNetowrk, clientAddr, serverAddr, &initMsg)
	if err != nil {
		fmt.Println("Erro:", err)
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
