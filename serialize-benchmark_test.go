package kosuzu_test

import (
	"testing"

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

	kosuzu.Serialize(18, choice)
}

func BenchmarkSerializeCustom(b *testing.B) {
	serialize := func(opcode int32, choice *Choice) *kosuzu.Packet {
		builder := kosuzu.NewPacketBuilder()

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

	serialize(18, choice)
}

func BenchmarkDeserializeReflect(b *testing.B) {
	choice := &Choice{
		Parameter:  34,
		Numbers:    []int64{13, 14, 15, 16, 18},
		Parameters: []complex128{2 + 3i, 3 + 1i, 2 + 5i},
	}

	packet, _ := kosuzu.Serialize(18, choice)
	restored := new(Choice)
	kosuzu.Deserialize(packet, &restored)
}

func BenchmarkDeserializeCustom(b *testing.B) {
	serialize := func(opcode int32, choice *Choice) *kosuzu.Packet {
		builder := kosuzu.NewPacketBuilder()

		builder.AddInt32(choice.Parameter)
		builder.AddInt64Array(choice.Numbers)

		return builder.BuildPacket(opcode)
	}
	deserialize := func(packet *kosuzu.Packet) *Choice {
		decomposer := kosuzu.NewPacketDecomposer(packet)
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
	deserialize(packet)
}
