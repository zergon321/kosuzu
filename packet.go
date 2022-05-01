package kosuzu

import (
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
	n, err := stream.Write(packet.data)

	if err != nil {
		return int64(n), err
	}

	return int64(n), nil
}

// ReadPacketFrom reads a new
// packet from the reader stream.
func ReadPacketFrom(stream io.Reader, order binary.ByteOrder, buffer []byte) (int64, Packet, error) {
	if order == nil {
		order = binary.BigEndian
	}

	packet := Packet{}

	// Read the opcode.
	var opcodeBuffer [4]byte
	n, err := stream.Read(opcodeBuffer[:])

	if err != nil {
		return int64(n), Packet{}, err
	}

	packet.Opcode = int32(order.Uint32(opcodeBuffer[:]))

	// Read the payload length.
	var sizeBuffer [8]byte
	n, err = stream.Read(sizeBuffer[:])

	if err != nil {
		return int64(n) + 4, Packet{}, err
	}

	packet.payloadLength = int64(order.Uint64(sizeBuffer[:]))

	// Read the payload.
	if buffer == nil {
		buffer = make([]byte, packet.payloadLength+12)
	}

	copy(buffer, opcodeBuffer[:])
	copy(buffer[4:], sizeBuffer[:])
	n, err = stream.Read(buffer[12:])

	if err != nil {
		return int64(n + 12), Packet{}, err
	}

	packet.data = buffer

	return int64(len(packet.data)), packet, nil
}

// PacketFromBytes creates a new packet
// out of the byte sequence.
func PacketFromBytes(data []byte, order binary.ByteOrder) (Packet, error) {
	if order == nil {
		order = binary.BigEndian
	}

	packet := Packet{
		Opcode:        int32(order.Uint32(data)),
		payloadLength: int64(order.Uint64(data[4:])),
		data:          data,
	}

	return packet, nil
}

// newPacket creates a new packet
// with the specified opcode. Opcodes
// are required to identify the type
// of the network packet.
func newPacket(opcode int32, data []byte, order binary.ByteOrder) Packet {
	payloadLength := int64(len(data)) - 12

	if order == nil {
		order = binary.BigEndian
	}

	order.PutUint32(data, uint32(opcode))
	order.PutUint64(data[4:], uint64(payloadLength))

	return Packet{
		Opcode:        opcode,
		payloadLength: payloadLength,
		data:          data,
	}
}

func NewPacketJS(opcode int32, payload []byte) Packet {
	var opcodeBuffer [4]byte
	binary.BigEndian.PutUint32(opcodeBuffer[:], uint32(opcode))

	var sizeBuffer [8]byte
	binary.BigEndian.PutUint64(sizeBuffer[:], uint64(len(payload)))

	data := make([]byte, 0, 12+len(payload))
	data = append(data, opcodeBuffer[:]...)
	data = append(data, sizeBuffer[:]...)
	data = append(data, payload...)

	return Packet{
		Opcode:        opcode,
		payloadLength: int64(len(payload)),
		data:          data,
	}
}
