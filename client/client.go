package main

import (
	"encoding/json"
	"fmt"
	"go-rcp/add"
	"log"
	"net"
)

func main() {
	res := rpcadd(10, 15)
	fmt.Println("Add result:", res)
}

func rpcadd(argA int64, argB int64) int64 {
	addr := "localhost:9000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Failed to dial")
		log.Fatal(err)
	}

	defer conn.Close()

	msg := createMsg(argA, argB)
	conn.Write(msg)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read message")
		log.Fatal(err)
	}

	res := add.ReadAddResult(buf[:n])
	return res.Result
}

func createMsg(argA int64, argB int64) []byte {
	msg := add.AddMsg{ArgA: argA, ArgB: argB}
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Failed to marshal message")
		log.Fatal(err)
	}

	return data
}
