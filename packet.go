package kosuzu

import (
	"bytes"
	"encoding/binary"
)

// Packet is a byte sequence
// representing a message
// serialized to be sent
// through network.
type Packet struct {
	Opcode int32
	Length int64
	Data   []byte
}

// NewPacket creates a new packet
// with the specified opcode. Opcodes
// are required to identify the type
// of the network packet.
func NewPacket(opcode int32, data []byte) *Packet {
	return &Packet{
		Opcode: opcode,
		Length: int64(len(data)),
		Data:   data,
	}
}

// Bytes returns the raw binary representation
// of the packet.
func (packet *Packet) Bytes() ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte,
		4+8+len(packet.Data)))
	err := binary.Write(buffer,
		binary.BigEndian, packet.Opcode)

	if err != nil {
		return nil, err
	}

	err = binary.Write(buffer,
		binary.BigEndian, packet.Length)

	if err != nil {
		return nil, err
	}

	_, err = buffer.Write(packet.Data)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
