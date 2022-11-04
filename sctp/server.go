package sctp

import (
	"fmt"
	"syscall"
	"unisinos/redes-i/tgb/address"
	"unisinos/redes-i/tgb/sniffer"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunServer(ip string, port int, responseAddresses []address.Address) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	SCTPAddr, err := sctp.MakeSCTPAddr(SCTPNetowrk, addr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	initMsg := NewSCTPInitMessage()
	listener, err := sctp.ListenSCTP(SCTPNetowrk, syscall.SOCK_STREAM, SCTPAddr, &initMsg)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Servidor executando no endereço %s\n", SCTPAddr.String())

	for {
		// Aguarda um conexão
		conn, err := listener.AcceptSCTP()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if remote := conn.RemoteAddr(); nil != remote {
			fmt.Println("Server side: conexão estabelecida com o endereço:", remote)
		}
		// obtém os dados recebidos do client
		go sniffAndSentToClient(conn)
	}
}

func sniffAndSentToClient(conn *sctp.SCTPConn) {
	remoteAddr := conn.RemoteAddr().String()
	for _, pkt := range sniffer.Sniff() {
		_, err := conn.SendMsg([]byte(pkt), nil)
		if err != nil {
			fmt.Printf("Server side: Erro -> %s\n", err)
		}
	}

	if err := conn.Close(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Server side: conexão encerrada com o endereço %s\n", remoteAddr)
	}
}
