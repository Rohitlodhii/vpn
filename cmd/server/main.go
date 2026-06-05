package main

import (
	"fmt"
	"myapp/protocol"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening")

	for {

		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	packet, _ := protocol.Decode(buf[:n])

	fmt.Println(string(packet.Message))

	resp := protocol.Packet{
		Type:    protocol.Pong,
		Message: []byte("pong"),
	}

	data, _ := protocol.Encode(resp)

	conn.Write(data)
}
