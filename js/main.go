package main

//go:generate gopherjs build -o kosuzu.js --minify

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/zergon321/kosuzu"
)

func serialize(opcode int32, obj map[string]interface{}) map[string]interface{} {
	builder := kosuzu.NewPacketBuilder()

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

		case []complex64:
			builder.AddComplex64Array(v)

		case []complex128:
			builder.AddComplex128Array(v)
		}
	}

	packet := builder.BuildPacket(opcode)
	jsPacket := map[string]interface{}{
		"opcode":     packet.Opcode,
		"dataLength": packet.DataLength(),
		"payload":    packet.Payload(),
	}

	return jsPacket
}

func main() {
	js.Module.Get("exports").Set("serialize", serialize)
}
