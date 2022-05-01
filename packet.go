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
	Opcode        int32
	payloadLength int64
	data          []byte
}

// Payload returns the data written
// to the network packet.
func (packet *Packet) Payload() []byte {
	payload := make([]byte, len(packet.data[12:]))
	copy(payload, packet.data)

	return payload
}

func (packet *Packet) PayloadNoCopy() []byte {
	return packet.data[12:]
}

// PayloadLength returns the length of
// the network packet payload.
func (packet *Packet) PayloadLength() int64 {
	return packet.payloadLength
}

// Bytes returns the raw binary representation
// of the packet.
func (packet *Packet) Bytes() ([]byte, error) {
	data := make([]byte, len(packet.data))
	copy(data, packet.data)

	return data, nil
}

func (packet *Packet) BytesNoCopy() ([]byte, error) {
	return packet.data, nil
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
		binary.BigEndian, packet.payloadLength)

	if err != nil {
		return 12, err
	}

	n, err := stream.Write(packet.data)

	if err != nil {
		return int64(n), err
	}

	return 12 + packet.payloadLength, nil
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
		binary.BigEndian, &packet.payloadLength)

	if err != nil {
		return 12, nil, err
	}

	packet.data = make([]byte, packet.payloadLength)
	n, err := stream.Read(packet.data)

	if err != nil {
		return int64(n), nil, err
	}

	return 12 + packet.payloadLength, packet, nil
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

// newPacket creates a new packet
// with the specified opcode. Opcodes
// are required to identify the type
// of the network packet.
func newPacket(opcode int32, data []byte, order binary.ByteOrder) *Packet {
	payloadLength := int64(len(data)) - 12

	if order == nil {
		order = binary.BigEndian
	}

	order.PutUint32(data, uint32(opcode))
	order.PutUint64(data[4:], uint64(payloadLength))

	return &Packet{
		Opcode:        opcode,
		payloadLength: payloadLength,
		data:          data,
	}
}
