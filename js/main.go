package main

//go:generate gopherjs build -o kosuzu.js --minify

import (
	"bytes"
	"encoding/binary"

	"github.com/gopherjs/gopherjs/js"
	"github.com/zergon321/kosuzu"
)

func packetBytes(packet map[string]interface{}) []byte {
	opcode := int32(packet["opcode"].(float64))
	dataLength := int64(packet["dataLength"].(float64))
	payload := packet["payload"].([]byte)

	data := make([]byte, 4+8+dataLength)
	buffer := bytes.NewBuffer(data)
	buffer.Reset()

	binary.Write(buffer, binary.BigEndian, opcode)
	binary.Write(buffer, binary.BigEndian, dataLength)
	buffer.Write(payload)

	return buffer.Bytes()
}

func packetFromBytes(data []byte) map[string]interface{} {
	packet, _ := kosuzu.PacketFromBytes(data, binary.BigEndian)
	jsPacket := map[string]interface{}{
		"opcode":        packet.Opcode,
		"payloadLength": packet.PayloadLength(),
		"payload":       packet.Payload(),
	}

	return jsPacket
}

func serialize(opcode int32, obj map[string]interface{}, scheme map[string]interface{}) map[string]interface{} {
	builder := kosuzu.NewPacketBuilder(0, binary.BigEndian)

	if len(scheme) > 0 {
		for param, value := range obj {
			switch scheme[param] {
			case "string":
				builder.AddString(value.(string))

			case "bool":
				builder.AddBool(value.(bool))

			case "byte":
				builder.AddByte(byte(value.(float64)))

			case "rune":
				builder.AddRune(value.(rune))

			case "int8":
				builder.AddInt8(int8(value.(float64)))

			case "uint8":
				builder.AddUint8(uint8(value.(float64)))

			case "int16":
				builder.AddInt16(int16(value.(float64)))

			case "uint16":
				builder.AddUint16(uint16(value.(float64)))

			case "int32":
				builder.AddInt32(int32(value.(float64)))

			case "uint32":
				builder.AddUint32(uint32(value.(float64)))

			case "int64":
				builder.AddInt64(int64(value.(float64)))

			case "uint64":
				builder.AddUint64(uint64(value.(float64)))

			case "float32":
				builder.AddFloat32(float32(value.(float64)))

			case "float64":
				builder.AddFloat64(value.(float64))

			case "[]bool":
				builder.AddBoolArray(value.([]bool))

			case "[]byte":
				builder.AddByteArray(value.([]byte))

			case "[]rune":
				builder.AddRuneArray(value.([]rune))

			case "[]int8":
				builder.AddInt8Array(value.([]int8))

			case "[]uint8":
				builder.AddUint8Array(value.([]uint8))

			case "[]int16":
				builder.AddInt16Array(value.([]int16))

			case "[]uint16":
				builder.AddUint16Array(value.([]uint16))

			case "[]int32":
				builder.AddInt32Array(value.([]int32))

			case "[]uint32":
				builder.AddUint32Array(value.([]uint32))

			case "[]int64":
				builder.AddInt64Array(value.([]int64))

			case "[]uint64":
				builder.AddUint64Array(value.([]uint64))

			case "[]float32":
				builder.AddFloat32Array(value.([]float32))

			case "[]float64":
				builder.AddFloat64Array(value.([]float64))
			}
		}
	} else {
		for _, value := range obj {
			switch v := value.(type) {
			case string:
				builder.AddString(v)

			case bool:
				builder.AddBool(v)

			case int8:
				builder.AddInt8(v)

			case uint8:
				builder.AddUint8(v)

			case int16:
				builder.AddInt16(v)

			case uint16:
				builder.AddUint16(v)

			case int32:
				builder.AddInt32(v)

			case uint32:
				builder.AddUint32(v)

			case int64:
				builder.AddInt64(v)

			case uint64:
				builder.AddUint64(v)

			case float32:
				builder.AddFloat32(v)

			case float64:
				builder.AddFloat64(v)

			case complex64:
				builder.AddComplex64(v)

			case complex128:
				builder.AddComplex128(v)

			case []bool:
				builder.AddBoolArray(v)

			case []int8:
				builder.AddInt8Array(v)

			case []uint8:
				builder.AddUint8Array(v)

			case []int16:
				builder.AddInt16Array(v)

			case []uint16:
				builder.AddUint16Array(v)

			case []int32:
				builder.AddInt32Array(v)

			case []uint32:
				builder.AddUint32Array(v)

			case []int64:
				builder.AddInt64Array(v)

			case []uint64:
				builder.AddUint64Array(v)

			case []float32:
				builder.AddFloat32Array(v)

			case []float64:
				builder.AddFloat64Array(v)
			}
		}
	}

	packet := builder.BuildPacket(opcode)
	jsPacket := map[string]interface{}{
		"opcode":        packet.Opcode,
		"payloadLength": packet.PayloadLength(),
		"payload":       packet.Payload(),
	}

	return jsPacket
}

func deserialize(scheme map[string]interface{}, packet map[string]interface{}) map[string]interface{} {
	goPacket := kosuzu.NewPacketJS(
		int32(packet["opcode"].(float64)),
		packet["payload"].([]byte))
	decomposer := kosuzu.NewPacketDecomposer(
		goPacket, binary.BigEndian)
	obj := map[string]interface{}{}

	for param, typ := range scheme {
		if typName, ok := typ.(string); ok {
			switch typName {
			case "string":
				val, _ := decomposer.ReadString()
				obj[param] = val

			case "bool":
				val, _ := decomposer.ReadBool()
				obj[param] = val

			case "byte":
				val, _ := decomposer.ReadByte()
				obj[param] = val

			case "rune":
				val, _ := decomposer.ReadRune()
				obj[param] = val

			case "int8":
				val, _ := decomposer.ReadInt8()
				obj[param] = val

			case "uint8":
				val, _ := decomposer.ReadUint8()
				obj[param] = val

			case "int16":
				val, _ := decomposer.ReadInt16()
				obj[param] = val

			case "uint16":
				val, _ := decomposer.ReadUint16()
				obj[param] = val

			case "int32":
				val, _ := decomposer.ReadInt32()
				obj[param] = val

			case "uint32":
				val, _ := decomposer.ReadUint32()
				obj[param] = val

			case "int64":
				val, _ := decomposer.ReadInt64()
				obj[param] = val

			case "uint64":
				val, _ := decomposer.ReadUint64()
				obj[param] = val

			case "float32":
				val, _ := decomposer.ReadFloat32()
				obj[param] = val

			case "float64":
				val, _ := decomposer.ReadFloat64()
				obj[param] = val

			case "[]bool":
				val, _ := decomposer.ReadBoolArray()
				obj[param] = val

			case "[]byte":
				val, _ := decomposer.ReadByteArray()
				obj[param] = val

			case "[]rune":
				val, _ := decomposer.ReadRuneArray()
				obj[param] = val

			case "[]int8":
				val, _ := decomposer.ReadInt8Array()
				obj[param] = val

			case "[]uint8":
				val, _ := decomposer.ReadUint8Array()
				obj[param] = val

			case "[]int16":
				val, _ := decomposer.ReadInt16Array()
				obj[param] = val

			case "[]uint16":
				val, _ := decomposer.ReadUint16Array()
				obj[param] = val

			case "[]int32":
				val, _ := decomposer.ReadInt32Array()
				obj[param] = val

			case "[]uint32":
				val, _ := decomposer.ReadUint32Array()
				obj[param] = val

			case "[]int64":
				val, _ := decomposer.ReadInt64Array()
				obj[param] = val

			case "[]uint64":
				val, _ := decomposer.ReadUint64Array()
				obj[param] = val

			case "[]float32":
				val, _ := decomposer.ReadFloat32Array()
				obj[param] = val

			case "[]float64":
				val, _ := decomposer.ReadFloat64Array()
				obj[param] = val
			}
		}
	}

	return obj
}

func main() {
	exports := map[string]interface{}{
		"serialize":       serialize,
		"deserialize":     deserialize,
		"packetBytes":     packetBytes,
		"packetFromBytes": packetFromBytes,
	}

	js.Global.Set("kosuzu", exports)
}
