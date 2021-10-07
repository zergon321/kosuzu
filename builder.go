package kosuzu

import (
	"bytes"
	"encoding/binary"
)

// Builder allows you to write values of different types into the packet.
type Builder struct {
	buffer *bytes.Buffer
}

// AddBool adds a bool value to the packet.
func (builder *Builder) AddBool(val bool) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddRune adds a rune value to the packet.
func (builder *Builder) AddRune(val rune) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddByte adds a byte value to the packet.
func (builder *Builder) AddByte(val byte) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddInt8 adds an int8 value to the packet.
func (builder *Builder) AddInt8(val int8) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddInt16 adds an int16 value to the packet.
func (builder *Builder) AddInt16(val int16) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddInt32 adds an int32 value to the packet.
func (builder *Builder) AddInt32(val int32) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddInt64 adds an int64 value to the packet.
func (builder *Builder) AddInt64(val int64) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddUint8 adds a uint8 value to the packet.
func (builder *Builder) AddUint8(val uint8) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddUint16 adds a uint16 value to the packet.
func (builder *Builder) AddUint16(val uint16) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddUint32 adds a uint32 value to the packet.
func (builder *Builder) AddUint32(val uint32) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddUint64 adds a uint64 value to the packet.
func (builder *Builder) AddUint64(val uint64) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddFloat32 adds a float32 value to the packet.
func (builder *Builder) AddFloat32(val float32) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddFloat64 adds a float64 value to the packet.
func (builder *Builder) AddFloat64(val float64) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddComplex64 adds a complex64 value to the packet.
func (builder *Builder) AddComplex64(val complex64) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddComplex128 adds a complex128 value to the packet.
func (builder *Builder) AddComplex128(val complex128) error {
	return binary.Write(builder.buffer, binary.BigEndian, val)
}

// AddString adds a string value to the packet.
func (builder *Builder) AddString(val string) error {
	// Write the string length.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	_, err = builder.buffer.Write([]byte(val))

	return err
}

// AddBytes adds bytes value to the packet.
func (builder *Builder) AddBytes(val []byte) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}
	_, err = builder.buffer.Write(val)

	return err
}

func (builder *Builder) AddInt8Array(val []int8) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddUint8Array(val []uint8) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddInt16Array(val []int16) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 2))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddUint16Array(val []uint16) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 2))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddInt32Array(val []int32) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 4))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddUint32Array(val []uint32) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 4))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddInt64Array(val []int64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 8))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddUint64Array(val []uint64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 8))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddFloat32Array(val []float32) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 4))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddFloat64Array(val []float64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 8))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddComplex64Array(val []complex64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 8))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddComplex128Array(val []complex128) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 16))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddBoolArray(val []bool) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

func (builder *Builder) AddRunes(val []rune) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val) * 4))

	if err != nil {
		return err
	}

	err = binary.Write(builder.buffer, binary.BigEndian, val)

	return err
}

// BuildPacket returns a packet with written values.
func (builder *Builder) BuildPacket(opcode int32) *Packet {
	return NewPacket(opcode, builder.buffer.Bytes())
}

// NewPacketBuilder creates a new packet builder
// to write values of certain types into the packet.
func NewPacketBuilder() *Builder {
	return &Builder{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
