package main

import (
	"fmt"
	"myapp/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	packet := protocol.Packet{
		Type:    protocol.Ping,
		Message: []byte("hello"),
	}

	data, err := protocol.Encode(packet)
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write(data); err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	resp, err := protocol.Decode(buf[:n])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp.Message))
}
