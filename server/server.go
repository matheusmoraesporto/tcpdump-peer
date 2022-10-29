package server

import (
	"encoding/hex"
	"fmt"
	"syscall"
	"unisinos/redes-i/tgb/common"

	sctp "github.com/thebagchi/sctp-go"
)

func main() {
	run("127.0.0.1", 12345)
}

func RunServer(ip string, port int) {
	run(ip, port)
}

func run(ip string, port int) {
	staddr := fmt.Sprintf("%s:%d", ip, port)
	addr, err := sctp.MakeSCTPAddr(common.SCTPNetowrk, staddr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	initMsg := common.NewSCTPInitMessage()
	server, err := sctp.ListenSCTP(common.SCTPNetowrk, syscall.SOCK_STREAM, addr, &initMsg)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer server.Close()

	for {
		// Aguarda um conexão
		conn, err := server.AcceptSCTP()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Conexão encontrada
		if remote := conn.RemoteAddr(); nil != remote {
			fmt.Println("Conexão estabelecida com o endereço:", remote)
		}
		// obtém os dados recebidos do client
		go handleClient(conn)
	}
}

func handleClient(conn *sctp.SCTPConn) {
	data := make([]byte, 8192)
	flag := 0

	for {
		info := &sctp.SCTPSndRcvInfo{}
		len, err := conn.RecvMsg(data, info, &flag)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		if len == 0 {
			fmt.Println("Conexão encerrada!")
			break
		} else {
			fmt.Println(fmt.Sprintf("Rcvd %d bytes", len))
			buffer := data[:len]
			fmt.Println(string(buffer))
			fmt.Println(hex.Dump(sctp.Pack(info)))
		}
	}
}
