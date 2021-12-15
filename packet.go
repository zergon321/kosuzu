package kosuzu

import (
	"bytes"
	"encoding/binary"
	"io"
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
	_, err := packet.WriteTo(buffer)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// WriteTo writes the whole contents of
// the packet to the writer stream.
func (packet *Packet) WriteTo(stream io.Writer) (int64, error) {
	err := binary.Write(stream,
		binary.BigEndian, packet.Opcode)

	if err != nil {
		return 4, err
	}

	err = binary.Write(stream,
		binary.BigEndian, packet.dataLength)

	if err != nil {
		return 12, err
	}

	n, err := stream.Write(packet.payload)

	if err != nil {
		return int64(n), err
	}

	return 12 + packet.dataLength, nil
}

// ReadPacketFrom reads a new
// packet from the reader stream.
func ReadPacketFrom(stream io.Reader) (int64, *Packet, error) {
	packet := new(Packet)

	err := binary.Read(stream,
		binary.BigEndian, &packet.Opcode)

	if err != nil {
		return 4, nil, err
	}

	err = binary.Read(stream,
		binary.BigEndian, &packet.dataLength)

	if err != nil {
		return 12, nil, err
	}

	packet.payload = make([]byte, packet.dataLength)
	n, err := stream.Read(packet.payload)

	if err != nil {
		return int64(n), nil, err
	}

	return 12 + packet.dataLength, packet, nil
}

// PacketFromBytes creates a new packet
// out of the byte sequence.
func PacketFromBytes(data []byte) (*Packet, error) {
	buffer := bytes.NewBuffer(data)
	_, packet, err := ReadPacketFrom(buffer)

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
