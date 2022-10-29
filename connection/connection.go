package connection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type Connection struct {
	Ip         string `json:"ip"`
	ServerPort int    `json:"serverPort"`
	ClientPort int    `json:"clientPort"`
}

func GetConnections(filePath string) (server Connection, clients []Connection, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("Erro: %v\n", err)
		return
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Arquivo de conexões mal formatado: %v\n", err)
		return
	}

	var connections []Connection
	err = json.Unmarshal(byteValue, &connections)
	if err != nil {
		return Connection{}, nil, err
	}

	return GetServerClients(connections)
}

func GetServerClients(connections []Connection) (server Connection, clients []Connection, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return Connection{}, nil, fmt.Errorf("Erro ao obter o hostname: %v", err)
	}

	localAddrs, err := net.LookupHost(hostname)
	if err != nil || len(localAddrs) < 1 {
		return Connection{}, nil, fmt.Errorf("Erro ao tentar identificar o endereço local: %v\n", err)
	}

	for _, c := range connections {
		if c.Ip == localAddrs[0] {
			server = c
		} else {
			clients = append(clients, c)
		}
	}

	return
}
