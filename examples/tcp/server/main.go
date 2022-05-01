package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/zergon321/kosuzu"
)

type PlayerMovement struct {
	ID int32
	X  float64
	Y  float64
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9828")
	handleError(err)
	conn, err := listener.Accept()
	handleError(err)

	_, packet, err := kosuzu.ReadPacketFrom(conn, binary.BigEndian, nil)
	handleError(err)
	var mvData PlayerMovement
	err = kosuzu.Deserialize(packet, &mvData, binary.BigEndian)
	handleError(err)

	fmt.Println(packet.Opcode)
	fmt.Println(mvData)

	err = conn.Close()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
