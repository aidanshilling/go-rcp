package main

import (
	"fmt"
	"go-rcp/add"
	"log"
	"net"
)

func main() {
	res1 := rpcadd(10, 15)
	res2 := rpcadd(10, 20)
	fmt.Println("Add result 1:", res1)
	fmt.Println("Add result 2:", res2)
}

func rpcadd(argA int64, argB int64) int64 {
	addr := "localhost:9000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Failed to dial")
		log.Fatal(err)
	}

	defer conn.Close()

	msg := add.MarshalMsg(argA, argB)
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
