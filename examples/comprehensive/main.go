package main

import (
	"fmt"

	"github.com/zergon321/kosuzu"
)

type Data struct {
	Str         string
	Rune        rune
	Runes       []rune
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int8        int8
	Uint8       uint8
	Int8s       []int8
	Uint8s      []uint8
	Int16       int16
	Uint16      uint16
	Int16s      []int16
	Uint16s     []uint16
	Int32       int32
	Uint32      uint32
	Int32s      []int32
	Uint32s     []uint32
	Int64       int64
	Uint64      uint64
	Int64s      []int64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
}

func main() {
	data := &Data{
		Str:         "Lolk",
		Rune:        '良',
		Runes:       []rune("ロックしましょう"),
		Bool:        true,
		Bools:       []bool{true, false, true, true, false},
		Byte:        32,
		Bytes:       []byte{192, 168, 1, 41},
		Int8:        -120,
		Uint8:       255,
		Int8s:       []int8{-120, 122, -125},
		Uint8s:      []uint8{1, 2, 3, 4, 5},
		Int16:       -120,
		Uint16:      255,
		Int16s:      []int16{-120, 122, -125},
		Uint16s:     []uint16{1, 2, 3, 4, 5},
		Int32:       -120,
		Uint32:      255,
		Int32s:      []int32{-120, 122, -125},
		Uint32s:     []uint32{1, 2, 3, 4, 5},
		Int64:       -120,
		Uint64:      255,
		Int64s:      []int64{-120, 122, -125},
		Uint64s:     []uint64{1, 2, 3, 4, 5},
		Float32:     -120.56,
		Float32s:    []float32{-130.1, -150.12, -14.8},
		Float64:     -120.56,
		Float64s:    []float64{-130.1, -150.12, -14.8},
		Complex64:   3 + 2i,
		Complex64s:  []complex64{3 + 2i, 2 + 3i, 7 + 11i},
		Complex128:  7 + 9i,
		Complex128s: []complex128{9 + 8i, 5 + 12i},
	}

	packet, err := kosuzu.Serialize(13, data)
	handleError(err)
	rawData, err := packet.Bytes()
	handleError(err)

	fmt.Println("Packet:", rawData)
	fmt.Println("Packet length:", len(rawData))
	fmt.Println("Packet contents length:", packet.DataLength())

	restored := new(Data)
	err = kosuzu.Deserialize(packet, &restored)
	handleError(err)

	fmt.Println(restored)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
