package kosuzu

import (
	"bytes"
	"encoding/binary"
)

// Packet represents a single network message.
// It has a byte code indicating the type of the message
// and a data payload in the form of a byte slice.
type Packet struct {
	Opcode int32
	Length int64
	Data   []byte
}

// NewPacket creates a new packet.
// It expects a byteCode for the type of message and
// a data parameter in the form of a byte slice.
func NewPacket(opcode int32, data []byte) *Packet {
	return &Packet{
		Opcode: opcode,
		Length: int64(len(data)),
		Data:   data,
	}
}

// Bytes returns the raw byte slice serialization of the packet.
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
