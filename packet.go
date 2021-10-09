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
	Opcode     int32
	dataLength int64
	payload    []byte
}

// Payload returns the data written
// to the network packet.
func (packet *Packet) Payload() []byte {
	payload := make([]byte, len(packet.payload))
	copy(payload, packet.payload)

	return payload
}

// DataLength returns the length of
// the network packet payload.
func (packet *Packet) DataLength() int64 {
	return packet.dataLength
}

// Bytes returns the raw binary representation
// of the packet.
func (packet *Packet) Bytes() ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte,
		4+8+packet.dataLength))
	buffer.Reset()
	err := binary.Write(buffer,
		binary.BigEndian, packet.Opcode)

	if err != nil {
		return nil, err
	}

	err = binary.Write(buffer,
		binary.BigEndian, packet.dataLength)

	if err != nil {
		return nil, err
	}

	_, err = buffer.Write(packet.payload)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// PacketFromBytes creates a new packet
// out of the byte sequence.
func PacketFromBytes(data []byte) (*Packet, error) {
	buffer := bytes.NewBuffer(data)
	packet := new(Packet)

	err := binary.Read(buffer,
		binary.BigEndian, &packet.Opcode)

	if err != nil {
		return nil, err
	}

	err = binary.Read(buffer,
		binary.BigEndian, &packet.dataLength)

	if err != nil {
		return nil, err
	}

	packet.payload = make([]byte, packet.dataLength)

	err = binary.Read(buffer,
		binary.BigEndian, &packet.payload)

	if err != nil {
		return nil, err
	}

	return packet, nil
}

// NewPacket creates a new packet
// with the specified opcode. Opcodes
// are required to identify the type
// of the network packet.
func NewPacket(opcode int32, data []byte) *Packet {
	return &Packet{
		Opcode:     opcode,
		dataLength: int64(len(data)),
		payload:    data,
	}
}
