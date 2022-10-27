package connections

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
