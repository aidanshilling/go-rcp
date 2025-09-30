package main

import (
	"fmt"
	"go-rcp/add"
	"log"
	"net"
)

func main() {
	addr := "localhost:9000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Printf("Server started on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(c net.Conn) {
			defer func() {
				fmt.Println("Closed connection at:", conn.RemoteAddr())
				c.Close()
			}()

			fmt.Println("New connection from:", c.RemoteAddr())

			buf := make([]byte, 1024)
			n, err := c.Read(buf)
			if err != nil {
				panic(err)
			}

			msg := add.ReadAddMsg(buf[:n])
			data := add.MarshalAddResult(msg)

			fmt.Println("Incoming:", string(buf[:n]))
			if _, err := c.Write(data); err != nil {
				panic(err)
			}
		}(conn)
	}
}
