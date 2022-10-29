package connections

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

func GetConnections(filePath string) (connections []Connection, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("Erro: %v\n", err)
		return
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Erro|Arquivo de conexões mal formatado: %v\n", err)
		return
	}

	json.Unmarshal(byteValue, &connections)
	return
}

func GetLocalIp() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("Erro ao obter o hostname: %v", err)
	}

	localAddrs, err := net.LookupHost(hostname)
	if err != nil || len(localAddrs) < 1 {
		return "", fmt.Errorf("Erro ao tentar identificar o endereço local: %v\n", err)
	}

	return localAddrs[0], nil
}
