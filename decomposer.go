package kosuzu

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// Decomposer allows you to read
// values of different types from the packet.
type Decomposer struct {
	buffer          []byte
	currentPosition int
	order           binary.ByteOrder
}

// ReadBool reads a bool value from the packet.
func (decomposer *Decomposer) ReadBool() (bool, error) {
	if decomposer.currentPosition >= len(decomposer.buffer) {
		return false, fmt.Errorf("buffer depleted")
	}

	result := decomposer.buffer[decomposer.currentPosition] != 0
	decomposer.currentPosition++

	return result, nil
}

// ReadRune reads a rune value from the packet.
func (decomposer *Decomposer) ReadRune() (rune, error) {
	if decomposer.currentPosition+4-1 >= len(decomposer.buffer) {
		return 'や', fmt.Errorf("buffer depleted")
	}

	result := rune(decomposer.order.Uint32(
		decomposer.buffer[decomposer.currentPosition:]))
	decomposer.currentPosition += 4

	return result, nil
}

// ReadByte reads a byte value from the packet.
func (decomposer *Decomposer) ReadByte() (byte, error) {
	if decomposer.currentPosition >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := decomposer.buffer[decomposer.currentPosition]
	decomposer.currentPosition++

	return result, nil
}

// ReadInt8 reads an int8 value from the packet.
func (decomposer *Decomposer) ReadInt8() (int8, error) {
	if decomposer.currentPosition >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := int8(decomposer.buffer[decomposer.currentPosition])
	decomposer.currentPosition++

	return result, nil
}

// ReadInt16 reads an int16 value from the packet.
func (decomposer *Decomposer) ReadInt16() (int16, error) {
	if decomposer.currentPosition+2-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := int16(decomposer.order.Uint16(
		decomposer.buffer[decomposer.currentPosition:]))
	decomposer.currentPosition += 2

	return result, nil
}

// ReadInt32 reads an int32 value from the packet.
func (decomposer *Decomposer) ReadInt32() (int32, error) {
	if decomposer.currentPosition+4-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := int32(decomposer.order.Uint32(
		decomposer.buffer[decomposer.currentPosition:]))
	decomposer.currentPosition += 4

	return result, nil
}

// ReadInt64 reads an int64 value from the packet.
func (decomposer *Decomposer) ReadInt64() (int64, error) {
	if decomposer.currentPosition+8-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := int64(decomposer.order.Uint64(
		decomposer.buffer[decomposer.currentPosition:]))
	decomposer.currentPosition += 8

	return result, nil
}

// ReadUint8 reads a uint8 value from the packet.
func (decomposer *Decomposer) ReadUint8() (uint8, error) {
	if decomposer.currentPosition >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := decomposer.buffer[decomposer.currentPosition]
	decomposer.currentPosition++

	return result, nil
}

// ReadUint16 reads a uint16 value from the packet.
func (decomposer *Decomposer) ReadUint16() (uint16, error) {
	if decomposer.currentPosition+2-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := decomposer.order.Uint16(
		decomposer.buffer[decomposer.currentPosition:])
	decomposer.currentPosition += 2

	return result, nil
}

// ReadUint32 reads a uint32 value from the packet.
func (decomposer *Decomposer) ReadUint32() (uint32, error) {
	if decomposer.currentPosition+4-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := decomposer.order.Uint32(
		decomposer.buffer[decomposer.currentPosition:])
	decomposer.currentPosition += 4

	return result, nil
}

// ReadUint64 reads a uint64 value from the packet.
func (decomposer *Decomposer) ReadUint64() (uint64, error) {
	if decomposer.currentPosition+8-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := decomposer.order.Uint64(
		decomposer.buffer[decomposer.currentPosition:])
	decomposer.currentPosition += 8

	return result, nil
}

// ReadFloat32 reads a float32 value from the packet.
func (decomposer *Decomposer) ReadFloat32() (float32, error) {
	if decomposer.currentPosition+4-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := math.Float32frombits(decomposer.order.Uint32(
		decomposer.buffer[decomposer.currentPosition:]))
	decomposer.currentPosition += 4

	return result, nil
}

// ReadFloat64 reads a float64 value from the packet.
func (decomposer *Decomposer) ReadFloat64() (float64, error) {
	if decomposer.currentPosition+8-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	result := math.Float64frombits(decomposer.order.Uint64(
		decomposer.buffer[decomposer.currentPosition:]))
	decomposer.currentPosition += 8

	return result, nil
}

// ReadComplex64 reads a complex64 value from the packet.
func (decomposer *Decomposer) ReadComplex64() (complex64, error) {
	if decomposer.currentPosition+8-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	realPart := math.Float32frombits(decomposer.order.Uint32(
		decomposer.buffer[decomposer.currentPosition:]))
	imagPart := math.Float32frombits(decomposer.order.Uint32(
		decomposer.buffer[decomposer.currentPosition+4:]))
	result := complex(realPart, imagPart)

	decomposer.currentPosition += 8

	return result, nil
}

// ReadComplex128 reads a complex128 value from the packet.
func (decomposer *Decomposer) ReadComplex128() (complex128, error) {
	if decomposer.currentPosition+16-1 >= len(decomposer.buffer) {
		return 0, fmt.Errorf("buffer depleted")
	}

	realPart := math.Float64frombits(decomposer.order.Uint64(
		decomposer.buffer[decomposer.currentPosition:]))
	imagPart := math.Float64frombits(decomposer.order.Uint64(
		decomposer.buffer[decomposer.currentPosition+4:]))
	result := complex(realPart, imagPart)

	decomposer.currentPosition += 16

	return result, nil
}

// ReadString copies a string value from the packet.
func (decomposer *Decomposer) ReadString() (string, error) {
	// Read the string length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return "", err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return "", fmt.Errorf("buffer depleted")
	}

	strBytes := make([]byte, length)
	copy(strBytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	return *(*string)(unsafe.Pointer(&strBytes)), nil
}

// ReadByteArray copies bytes from the packet.
func (decomposer *Decomposer) ReadByteArray() ([]byte, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	return bytes, nil
}

func (decomposer *Decomposer) ReadInt8Array() ([]int8, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	data := *(*[]int8)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadUint8Array() ([]uint8, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	data := *(*[]uint8)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadInt16Array() ([]int16, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 2
	header.Cap /= 2
	data := *(*[]int16)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadUint16Array() ([]uint16, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 2
	header.Cap /= 2
	data := *(*[]uint16)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadInt32Array() ([]int32, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 4
	header.Cap /= 4
	data := *(*[]int32)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadUint32Array() ([]uint32, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 4
	header.Cap /= 4
	data := *(*[]uint32)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadInt64Array() ([]int64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 8
	header.Cap /= 8
	data := *(*[]int64)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadUint64Array() ([]uint64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 8
	header.Cap /= 8
	data := *(*[]uint64)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadFloat32Array() ([]float32, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 4
	header.Cap /= 4
	data := *(*[]float32)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadFloat64Array() ([]float64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 8
	header.Cap /= 8
	data := *(*[]float64)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadComplex64Array() ([]complex64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 8
	header.Cap /= 8
	data := *(*[]complex64)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadComplex128Array() ([]complex128, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 16
	header.Cap /= 16
	data := *(*[]complex128)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadBoolArray() ([]bool, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	data := *(*[]bool)(unsafe.Pointer(&header))

	return data, nil
}

func (decomposer *Decomposer) ReadRuneArray() ([]rune, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	if decomposer.currentPosition+int(length)-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, length)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+int(length)])
	decomposer.currentPosition += int(length)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Len /= 4
	header.Cap /= 4
	data := *(*[]rune)(unsafe.Pointer(&header))

	return data, nil
}

// ReadNBytes reads n bytes from the packet.
func (decomposer *Decomposer) ReadNBytes(n int) ([]byte, error) {
	if decomposer.currentPosition+n-1 >= len(decomposer.buffer) {
		return nil, fmt.Errorf("buffer depleted")
	}

	bytes := make([]byte, n)
	copy(bytes, decomposer.buffer[decomposer.
		currentPosition:decomposer.currentPosition+n])
	decomposer.currentPosition += n

	return bytes, nil
}

// NewPacketDecomposer creates a new packet decomposer
// to read values of certain types from the packet.
func NewPacketDecomposer(packet Packet, order binary.ByteOrder) *Decomposer {
	decomposer := &Decomposer{
		buffer:          packet.data,
		currentPosition: 12,
	}

	if order == nil {
		order = binary.BigEndian
	}

	decomposer.order = order

	return decomposer
}
