package kosuzu

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

// Builder allows you to write values of different types into the packet.
type Builder struct {
	buffer          []byte
	currentPosition int
	order           binary.ByteOrder
}

func (builder *Builder) growIfNeeded(size int) {
	if len(builder.buffer)-builder.currentPosition-1 < size {
		builder.buffer = append(builder.buffer, make([]byte, size)...)
	}

	builder.buffer = builder.buffer[:cap(builder.buffer)]
}

// AddBytes adds the byte sequence to
// the buffer without writing its size.
func (builder *Builder) AddBytes(val []byte) error {
	builder.growIfNeeded(len(val))
	copy(builder.buffer[builder.currentPosition:], val)
	builder.currentPosition += len(val)

	return nil
}

// AddBool adds a bool value to the packet.
func (builder *Builder) AddBool(val bool) error {
	builder.growIfNeeded(1)

	if val {
		builder.buffer[builder.currentPosition] = 1
	} else {
		builder.buffer[builder.currentPosition] = 0
	}

	builder.currentPosition++

	return nil
}

// AddRune adds a rune value to the packet.
func (builder *Builder) AddRune(val rune) error {
	builder.growIfNeeded(4)
	builder.order.PutUint32(builder.
		buffer[builder.currentPosition:], uint32(val))
	builder.currentPosition += 4

	return nil
}

// AddByte adds a byte value to the packet.
func (builder *Builder) AddByte(val byte) error {
	builder.growIfNeeded(1)
	builder.buffer[builder.currentPosition] = val
	builder.currentPosition++

	return nil
}

// AddInt8 adds an int8 value to the packet.
func (builder *Builder) AddInt8(val int8) error {
	builder.growIfNeeded(1)
	builder.buffer[builder.currentPosition] = byte(val)
	builder.currentPosition++

	return nil
}

// AddInt16 adds an int16 value to the packet.
func (builder *Builder) AddInt16(val int16) error {
	builder.growIfNeeded(2)
	builder.order.PutUint16(builder.
		buffer[builder.currentPosition:], uint16(val))
	builder.currentPosition += 2

	return nil
}

// AddInt32 adds an int32 value to the packet.
func (builder *Builder) AddInt32(val int32) error {
	builder.growIfNeeded(4)
	builder.order.PutUint32(builder.
		buffer[builder.currentPosition:], uint32(val))
	builder.currentPosition += 4

	return nil
}

// AddInt64 adds an int64 value to the packet.
func (builder *Builder) AddInt64(val int64) error {
	builder.growIfNeeded(8)
	builder.order.PutUint64(builder.
		buffer[builder.currentPosition:], uint64(val))
	builder.currentPosition += 8

	return nil
}

// AddUint8 adds a uint8 value to the packet.
func (builder *Builder) AddUint8(val uint8) error {
	builder.growIfNeeded(1)
	builder.buffer[builder.currentPosition] = val
	builder.currentPosition++

	return nil
}

// AddUint16 adds a uint16 value to the packet.
func (builder *Builder) AddUint16(val uint16) error {
	builder.growIfNeeded(2)
	builder.order.PutUint16(builder.
		buffer[builder.currentPosition:], val)
	builder.currentPosition += 2

	return nil
}

// AddUint32 adds a uint32 value to the packet.
func (builder *Builder) AddUint32(val uint32) error {
	builder.growIfNeeded(4)
	builder.order.PutUint32(builder.
		buffer[builder.currentPosition:], val)
	builder.currentPosition += 4

	return nil
}

// AddUint64 adds a uint64 value to the packet.
func (builder *Builder) AddUint64(val uint64) error {
	builder.growIfNeeded(8)
	builder.order.PutUint64(builder.
		buffer[builder.currentPosition:], val)
	builder.currentPosition += 8

	return nil
}

// AddFloat32 adds a float32 value to the packet.
func (builder *Builder) AddFloat32(val float32) error {
	builder.growIfNeeded(4)
	builder.order.PutUint32(builder.
		buffer[builder.currentPosition:], math.Float32bits(val))
	builder.currentPosition += 4

	return nil
}

// AddFloat64 adds a float64 value to the packet.
func (builder *Builder) AddFloat64(val float64) error {
	builder.growIfNeeded(8)
	builder.order.PutUint64(builder.
		buffer[builder.currentPosition:], math.Float64bits(val))
	builder.currentPosition += 8

	return nil
}

// AddComplex64 adds a complex64 value to the packet.
func (builder *Builder) AddComplex64(val complex64) error {
	builder.growIfNeeded(8)

	realPart := float32(real(val))
	builder.order.PutUint32(builder.
		buffer[builder.currentPosition:], math.Float32bits(realPart))

	imagPart := float32(imag(val))
	builder.order.PutUint32(builder.
		buffer[builder.currentPosition+4:], math.Float32bits(imagPart))

	builder.currentPosition += 8

	return nil
}

// AddComplex128 adds a complex128 value to the packet.
func (builder *Builder) AddComplex128(val complex128) error {
	builder.growIfNeeded(16)

	builder.order.PutUint64(builder.
		buffer[builder.currentPosition:], math.Float64bits(real(val)))
	builder.order.PutUint64(builder.
		buffer[builder.currentPosition+8:], math.Float64bits(imag(val)))

	builder.currentPosition += 16

	return nil
}

// AddString adds a string value to the packet.
func (builder *Builder) AddString(val string) error {
	// Write the string length.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	strBytes := (*[0x7fff0000]byte)(unsafe.Pointer(
		(*reflect.StringHeader)(unsafe.Pointer(&val)).Data),
	)[:len(val):len(val)]
	err = builder.AddBytes(strBytes)

	if err != nil {
		return err
	}

	return nil
}

// AddByteArray adds bytes value to the packet.
func (builder *Builder) AddByteArray(val []byte) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	err = builder.AddBytes(val)

	return err
}

func (builder *Builder) AddInt8Array(val []int8) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddUint8Array(val []uint8) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddInt16Array(val []int16) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 2)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 2
	header.Cap *= 2
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddUint16Array(val []uint16) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 2)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 2
	header.Cap *= 2
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddInt32Array(val []int32) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 4)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 4
	header.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddUint32Array(val []uint32) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 4)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 4
	header.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddInt64Array(val []int64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 8)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 8
	header.Cap *= 8
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddUint64Array(val []uint64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 8)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 8
	header.Cap *= 8
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddFloat32Array(val []float32) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 4)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 4
	header.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddFloat64Array(val []float64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 8)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 8
	header.Cap *= 8
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddComplex64Array(val []complex64) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 8)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 8
	header.Cap *= 8
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddComplex128Array(val []complex128) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 16)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 16
	header.Cap *= 16
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddBoolArray(val []bool) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)))

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

func (builder *Builder) AddRuneArray(val []rune) error {
	// Write the number of bytes.
	err := builder.AddInt32(int32(len(val)) * 4)

	if err != nil {
		return err
	}

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&val))
	header.Len *= 4
	header.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(&header))
	err = builder.AddBytes(data)

	return err
}

// BuildPacket returns a packet with written values.
func (builder *Builder) BuildPacket(opcode int32) Packet {
	return newPacket(opcode,
		builder.buffer[:builder.
			currentPosition], builder.order)
}

// NewPacketBuilder creates a new packet builder
// to write values of certain types into the packet.
func NewPacketBuilder(initialLength int, order binary.ByteOrder) Builder {
	builder := Builder{
		buffer:          make([]byte, initialLength+12),
		currentPosition: 12,
	}

	if order == nil {
		order = binary.BigEndian
	}

	builder.order = order

	return builder
}
