package main

import (
	"net"

	"github.com/zergon321/kosuzu"
)

type PlayerMovement struct {
	ID int32
	X  float64
	Y  float64
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9828")
	handleError(err)

	opcode := int32(32)
	mvData := PlayerMovement{
		ID: 132,
		X:  116.198,
		Y:  20.07,
	}

	packet, err := kosuzu.Serialize(opcode, mvData)
	handleError(err)
	_, err = packet.WriteTo(conn)
	handleError(err)

	err = conn.Close()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
