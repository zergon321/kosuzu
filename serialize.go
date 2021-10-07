package kosuzu

import (
	"fmt"
	"reflect"
)

func readFromPacket(decomposer *Decomposer, fieldVal *reflect.Value, fieldTyp reflect.Type) error {
	switch fieldTyp.Kind() {
	case reflect.Int8:
		val, err := decomposer.ReadInt8()

		if err != nil {
			return err
		}

		fieldVal.SetInt(int64(val))

	case reflect.Uint8:
		val, err := decomposer.ReadUint8()

		if err != nil {
			return err
		}

		fieldVal.SetUint(uint64(val))

	case reflect.Int16:
		val, err := decomposer.ReadInt16()

		if err != nil {
			return err
		}

		fieldVal.SetInt(int64(val))

	case reflect.Uint16:
		val, err := decomposer.ReadUint16()

		if err != nil {
			return err
		}

		fieldVal.SetUint(uint64(val))

	case reflect.Int32:
		val, err := decomposer.ReadInt32()

		if err != nil {
			return err
		}

		fieldVal.SetInt(int64(val))

	case reflect.Uint32:
		val, err := decomposer.ReadUint32()

		if err != nil {
			return err
		}

		fieldVal.SetUint(uint64(val))

	case reflect.Int64:
		val, err := decomposer.ReadInt64()

		if err != nil {
			return err
		}

		fieldVal.SetInt(int64(val))

	case reflect.Uint64:
		val, err := decomposer.ReadUint64()

		if err != nil {
			return err
		}

		fieldVal.SetUint(uint64(val))

	case reflect.Float32:
		val, err := decomposer.ReadFloat32()

		if err != nil {
			return err
		}

		fieldVal.SetFloat(float64(val))

	case reflect.Float64:
		val, err := decomposer.ReadFloat64()

		if err != nil {
			return err
		}

		fieldVal.SetFloat(float64(val))

	case reflect.Complex64:
		val, err := decomposer.ReadComplex64()

		if err != nil {
			return err
		}

		fieldVal.SetComplex(complex128(val))

	case reflect.Complex128:
		val, err := decomposer.ReadComplex128()

		if err != nil {
			return err
		}

		fieldVal.SetComplex(complex128(val))

	case reflect.String:
		val, err := decomposer.ReadString()

		if err != nil {
			return err
		}

		fieldVal.SetString(val)

	case reflect.Bool:
		val, err := decomposer.ReadBool()

		if err != nil {
			return err
		}

		fieldVal.SetBool(val)

	case reflect.Slice:
		switch fieldVal.Interface().(type) {
		case []int8:
			slice, err := decomposer.ReadInt8Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []uint8:
			slice, err := decomposer.ReadUint8Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []int16:
			slice, err := decomposer.ReadInt16Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []uint16:
			slice, err := decomposer.ReadUint16Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []int32:
			slice, err := decomposer.ReadInt32Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []uint32:
			slice, err := decomposer.ReadUint32Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []int64:
			slice, err := decomposer.ReadInt64Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []uint64:
			slice, err := decomposer.ReadUint64Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []float32:
			slice, err := decomposer.ReadFloat32Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []float64:
			slice, err := decomposer.ReadFloat64Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []complex64:
			slice, err := decomposer.ReadComplex64Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []complex128:
			slice, err := decomposer.ReadComplex128Array()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		case []bool:
			slice, err := decomposer.ReadBoolArray()

			if err != nil {
				return err
			}

			fieldVal.Set(reflect.ValueOf(slice))

		default:
			return fmt.Errorf(
				"slice type is not supported: %v", fieldTyp)
		}

	default:
		return fmt.Errorf(
			"the field type is unsupported: %s", fieldTyp.Kind())
	}

	return nil
}

func writeToPacket(builder *Builder, fieldVal reflect.Value, fieldTyp reflect.Type) error {
	switch fieldTyp.Kind() {
	case reflect.Int8:
		err := builder.AddInt8(int8(fieldVal.Int()))

		if err != nil {
			return err
		}

	case reflect.Uint8:
		err := builder.AddUint8(uint8(fieldVal.Uint()))

		if err != nil {
			return err
		}

	case reflect.Int16:
		err := builder.AddInt16(int16(fieldVal.Int()))

		if err != nil {
			return err
		}

	case reflect.Uint16:
		err := builder.AddUint16(uint16(fieldVal.Uint()))

		if err != nil {
			return err
		}

	case reflect.Int32:
		err := builder.AddInt32(int32(fieldVal.Int()))

		if err != nil {
			return err
		}

	case reflect.Uint32:
		err := builder.AddUint32(uint32(fieldVal.Uint()))

		if err != nil {
			return err
		}

	case reflect.Int64:
		err := builder.AddInt64(fieldVal.Int())

		if err != nil {
			return err
		}

	case reflect.Uint64:
		err := builder.AddUint64(uint64(fieldVal.Uint()))

		if err != nil {
			return err
		}

	case reflect.Float32:
		err := builder.AddFloat32(float32(fieldVal.Float()))

		if err != nil {
			return err
		}

	case reflect.Float64:
		err := builder.AddFloat64(float64(fieldVal.Float()))

		if err != nil {
			return err
		}

	case reflect.Complex64:
		err := builder.AddComplex64(complex64(fieldVal.Complex()))

		if err != nil {
			return err
		}

	case reflect.Complex128:
		err := builder.AddComplex128(complex128(fieldVal.Complex()))

		if err != nil {
			return err
		}

	case reflect.String:
		err := builder.AddString(fieldVal.String())

		if err != nil {
			return err
		}

	case reflect.Bool:
		err := builder.AddBool(fieldVal.Bool())

		if err != nil {
			return err
		}

	case reflect.Slice:
		switch slice := fieldVal.Interface().(type) {
		case []int8:
			err := builder.AddInt8Array(slice)

			if err != nil {
				return err
			}

		case []uint8:
			err := builder.AddUint8Array(slice)

			if err != nil {
				return err
			}

		case []int16:
			err := builder.AddInt16Array(slice)

			if err != nil {
				return err
			}

		case []uint16:
			err := builder.AddUint16Array(slice)

			if err != nil {
				return err
			}

		case []int32:
			err := builder.AddInt32Array(slice)

			if err != nil {
				return err
			}

		case []uint32:
			err := builder.AddUint32Array(slice)

			if err != nil {
				return err
			}

		case []int64:
			err := builder.AddInt64Array(slice)

			if err != nil {
				return err
			}

		case []uint64:
			err := builder.AddUint64Array(slice)

			if err != nil {
				return err
			}

		case []float32:
			err := builder.AddFloat32Array(slice)

			if err != nil {
				return err
			}

		case []float64:
			err := builder.AddFloat64Array(slice)

			if err != nil {
				return err
			}

		case []complex64:
			err := builder.AddComplex64Array(slice)

			if err != nil {
				return err
			}

		case []complex128:
			err := builder.AddComplex128Array(slice)

			if err != nil {
				return err
			}

		case []bool:
			err := builder.AddBoolArray(slice)

			if err != nil {
				return err
			}

		default:
			return fmt.Errorf(
				"slice type is not supported: %v", fieldTyp)
		}

	default:
		return fmt.Errorf(
			"the field type is unsupported: %s", fieldTyp.Kind())
	}

	return nil
}

// Serialize serializes the given object
// and creates a network packet from it.
func Serialize(opcode int32, value interface{}) (*Packet, error) {
	builder := NewPacketBuilder()
	val := reflect.ValueOf(value)

	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			fieldVal := val.Field(i)
			fieldTyp := val.Field(i).Type()

			err := writeToPacket(builder,
				fieldVal, fieldTyp)

			if err != nil {
				return nil, err
			}
		}
	} else {
		err := writeToPacket(builder,
			val, val.Type())

		if err != nil {
			return nil, err
		}
	}

	return builder.BuildPacket(opcode), nil
}

// Deserialize deserializes the packet
// into the given object.
func Deserialize(packet *Packet, obj interface{}) error {
	decomposer := NewPacketDecomposer(packet)
	val := reflect.ValueOf(obj)

	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			fieldVal := val.Field(i)
			fieldTyp := val.Field(i).Type()

			err := readFromPacket(decomposer,
				&fieldVal, fieldTyp)

			if err != nil {
				return err
			}
		}
	} else {
		err := readFromPacket(decomposer,
			&val, val.Type())

		if err != nil {
			return err
		}
	}

	return nil
}
