package kosuzu

import (
	"bytes"
	"encoding/binary"
)

// Decomposer allows you to read
// values of different types from the packet.
type Decomposer struct {
	buffer *bytes.Reader
}

// ReadBool reads a bool value from the packet.
func (decomposer *Decomposer) ReadBool() (bool, error) {
	var result bool
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return false, err
	}

	return result, nil
}

// ReadRune reads a rune value from the packet.
func (decomposer *Decomposer) ReadRune() (rune, int, error) {
	var result rune
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 'や', 3, err
	}

	return result, len(string(result)), nil
}

// ReadByte reads a byte value from the packet.
func (decomposer *Decomposer) ReadByte() (byte, error) {
	var result byte
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadInt8 reads an int8 value from the packet.
func (decomposer *Decomposer) ReadInt8() (int8, error) {
	var result int8
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadInt16 reads an int16 value from the packet.
func (decomposer *Decomposer) ReadInt16() (int16, error) {
	var result int16
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadInt32 reads an int32 value from the packet.
func (decomposer *Decomposer) ReadInt32() (int32, error) {
	var result int32
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadInt64 reads an int64 value from the packet.
func (decomposer *Decomposer) ReadInt64() (int64, error) {
	var result int64
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadUint8 reads a uint8 value from the packet.
func (decomposer *Decomposer) ReadUint8() (uint8, error) {
	var result uint8
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadUint16 reads a uint16 value from the packet.
func (decomposer *Decomposer) ReadUint16() (uint16, error) {
	var result uint16
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadUint32 reads a uint32 value from the packet.
func (decomposer *Decomposer) ReadUint32() (uint32, error) {
	var result uint32
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadUint64 reads a uint64 value from the packet.
func (decomposer *Decomposer) ReadUint64() (uint64, error) {
	var result uint64
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadFloat32 reads a float32 value from the packet.
func (decomposer *Decomposer) ReadFloat32() (float32, error) {
	var result float32
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadFloat64 reads a float64 value from the packet.
func (decomposer *Decomposer) ReadFloat64() (float64, error) {
	var result float64
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadComplex64 reads a complex64 value from the packet.
func (decomposer *Decomposer) ReadComplex64() (complex64, error) {
	var result complex64
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadComplex128 reads a complex128 value from the packet.
func (decomposer *Decomposer) ReadComplex128() (complex128, error) {
	var result complex128
	err := binary.Read(decomposer.buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// ReadString reads a string value from the packet.
func (decomposer *Decomposer) ReadString() (string, error) {
	// Read the string length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length)
	_, err = decomposer.buffer.Read(strBytes)

	if err != nil {
		return "", err
	}

	return string(strBytes), nil
}

// ReadByteArray reads bytes from the packet.
func (decomposer *Decomposer) ReadByteArray() ([]byte, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	bytes := make([]byte, length)
	_, err = decomposer.buffer.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (decomposer *Decomposer) ReadInt8Array() ([]int8, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]int8, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadUint8Array() ([]uint8, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]uint8, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadInt16Array() ([]int16, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]int16, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadUint16Array() ([]uint16, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]uint16, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadInt32Array() ([]int32, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]int32, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadUint32Array() ([]uint32, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]uint32, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadInt64Array() ([]int64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]int64, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadUint64Array() ([]uint64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]uint64, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadFloat32Array() ([]float32, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]float32, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadFloat64Array() ([]float64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]float64, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadComplex64Array() ([]complex64, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]complex64, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadComplex128Array() ([]complex128, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]complex128, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadBoolArray() ([]bool, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]bool, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (decomposer *Decomposer) ReadRuneArray() ([]rune, error) {
	// Read the byte slice length.
	length, err := decomposer.ReadInt32()

	if err != nil {
		return nil, err
	}

	val := make([]rune, length)
	err = binary.Read(decomposer.buffer, binary.BigEndian, &val)

	if err != nil {
		return nil, err
	}

	return val, nil
}

// ReadNBytes reads n bytes from the packet.
func (decomposer *Decomposer) ReadNBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := decomposer.buffer.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// NewPacketDecomposer creates a new packet decomposer
// to read values of certain types from the packet.
func NewPacketDecomposer(packet *Packet) *Decomposer {
	return &Decomposer{
		buffer: bytes.NewReader(packet.data),
	}
}
