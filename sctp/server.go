package sctp

import (
	"fmt"
	"syscall"
	"unisinos/redes-i/tgb/address"
	"unisinos/redes-i/tgb/sniffer"

	sctp "github.com/thebagchi/sctp-go"
)

func (_ ConnectionSCTP) RunServer(ip string, port int, responseAddresses []address.Address) {
	staddr := fmt.Sprintf("%s:%d", ip, port)
	addr, err := sctp.MakeSCTPAddr(SCTPNetowrk, staddr)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	initMsg := NewSCTPInitMessage()
	server, err := sctp.ListenSCTP(SCTPNetowrk, syscall.SOCK_STREAM, addr, &initMsg)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer server.Close()

	for {
		fmt.Printf("Servidor executando no endereço %s\n", addr.String())
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
