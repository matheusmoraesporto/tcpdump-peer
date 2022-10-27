package main

import (
	"fmt"
	"unisinos/redes-i/tgb/connections"
)

func main() {
	connList, err := connections.GetConnections("./connections/addresses.json")
	if err != nil {
		return
	}

	fmt.Println(connList)
}
