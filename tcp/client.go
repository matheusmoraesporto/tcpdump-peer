package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func (_ ConnectionTCP) RunClient(ipLocal, ipRemote string, portLocal, portRemote int) {
	localAddr := HandleTCPAddress(ipLocal, 0)
	remoteAddr := HandleTCPAddress(ipRemote, portRemote)
	connection, err := net.DialTCP("tcp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// lendo entrada do terminal
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("texto a ser enviado: ")
		texto, erro2 := leitor.ReadString('\n')
		if erro2 != nil {
			fmt.Println(erro2)
		}

		// escrevendo a mensagem na conexão (socket)
		fmt.Fprintf(connection, texto+"\n")

		// ouvindo a resposta do servidor (eco)
		bufio.NewReader(connection).ReadString('\n')
		// if err3 != nil {
		// 	fmt.Println(err3)
		// }
		// // escrevendo a resposta do servidor no terminal
		// fmt.Print("Resposta do servidor: " + mensagem)
	}
}
