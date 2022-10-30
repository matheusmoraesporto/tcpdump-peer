package address

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type Address struct {
	Ip         string `json:"ip"`
	ServerPort int    `json:"serverPort"`
	ClientPort int    `json:"clientPort"`
}

func GetConnections(filePath string) (server Address, clients []Address, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("Erro: %v\n", err)
		return
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Arquivo de endereços mal formatado: %v\n", err)
		return
	}

	var addrs []Address
	err = json.Unmarshal(byteValue, &addrs)
	if err != nil {
		return Address{}, nil, err
	}

	return GetServerClients(addrs)
}

func GetServerClients(addrs []Address) (server Address, clients []Address, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return Address{}, nil, fmt.Errorf("Erro ao obter o hostname: %v", err)
	}

	localAddrs, err := net.LookupHost(hostname)
	if err != nil || len(localAddrs) < 1 {
		return Address{}, nil, fmt.Errorf("Erro ao tentar identificar o endereço local: %v\n", err)
	}

	for _, a := range addrs {
		if a.Ip == localAddrs[0] {
			server = a
		} else {
			clients = append(clients, a)
		}
	}

	return
}
