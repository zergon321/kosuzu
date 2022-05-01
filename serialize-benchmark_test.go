package kosuzu_test

import (
	"encoding/binary"
	"testing"
	"unsafe"

	"github.com/zergon321/kosuzu"
)

type Choice struct {
	Parameter  int32
	Numbers    []int64
	Parameters []complex128
}

func BenchmarkSerializeReflect(b *testing.B) {
	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	for i := 0; i < b.N; i++ {
		kosuzu.Serialize(18, choice, binary.BigEndian)
	}
}

func BenchmarkSerializeCustom(b *testing.B) {
	order := binary.BigEndian
	builder := kosuzu.NewPacketBuilder(512, order)

	serialize := func(opcode int32, choice *Choice) kosuzu.Packet {
		builder.Reset()

		builder.AddInt32(choice.Parameter)
		builder.AddInt64Array(choice.Numbers)
		builder.AddComplex128Array(choice.Parameters)

		return builder.BuildPacket(opcode)
	}

	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	for i := 0; i < b.N; i++ {
		serialize(18, choice)
	}
}

func BenchmarkSerializeCustomReuseBuffers(b *testing.B) {
	order := binary.BigEndian
	var buffer [512]byte
	builder, _ := kosuzu.NewPacketBuilderWithBuffer(buffer[:], order)

	serialize := func(opcode int32, choice *Choice) kosuzu.Packet {
		builder.ResetWithBuffer(buffer[:])

		builder.AddInt32(choice.Parameter)
		builder.AddInt64Array(choice.Numbers)
		builder.AddComplex128Array(choice.Parameters)

		return builder.BuildPacket(opcode)
	}

	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	for i := 0; i < b.N; i++ {
		serialize(18, choice)
	}
}

func BenchmarkDeserializeReflect(b *testing.B) {
	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	packet, _ := kosuzu.Serialize(18, choice, binary.BigEndian)
	restored := new(Choice)

	for i := 0; i < b.N; i++ {
		kosuzu.Deserialize(packet, &restored, binary.BigEndian)
	}
}

func BenchmarkDeserializeCustom(b *testing.B) {
	serialize := func(opcode int32, choice *Choice) kosuzu.Packet {
		builder := kosuzu.NewPacketBuilder(int(unsafe.Sizeof(Choice{})), binary.BigEndian)

		builder.AddInt32(choice.Parameter)
		builder.AddInt64Array(choice.Numbers)

		return builder.BuildPacket(opcode)
	}
	deserialize := func(packet kosuzu.Packet) *Choice {
		decomposer := kosuzu.NewPacketDecomposer(packet, binary.BigEndian)
		choice := new(Choice)

		choice.Parameter, _ = decomposer.ReadInt32()
		choice.Numbers, _ = decomposer.ReadInt64Array()
		choice.Parameters, _ = decomposer.ReadComplex128Array()

		return choice
	}

	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	packet := serialize(18, choice)

	for i := 0; i < b.N; i++ {
		deserialize(packet)
	}
}

func BenchmarkDeserializeCustomNoCopy(b *testing.B) {
	serialize := func(opcode int32, choice *Choice) kosuzu.Packet {
		builder := kosuzu.NewPacketBuilder(int(unsafe.Sizeof(Choice{})), binary.BigEndian)

		builder.AddInt32(choice.Parameter)
		builder.AddInt64Array(choice.Numbers)
		builder.AddComplex128Array(choice.Parameters)

		return builder.BuildPacket(opcode)
	}
	deserialize := func(packet kosuzu.Packet) Choice {
		var choice Choice
		decomposer := kosuzu.NewPacketDecomposer(packet, binary.BigEndian)

		choice.Parameter, _ = decomposer.ReadInt32()
		choice.Numbers, _ = decomposer.ReadInt64ArrayNoCopy()
		choice.Parameters, _ = decomposer.ReadComplex128ArrayNoCopy()

		return choice
	}

	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	packet := serialize(18, choice)

	for i := 0; i < b.N; i++ {
		deserialize(packet)
	}
}
