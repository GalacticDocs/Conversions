package reader

import (
	"bytes"
	"strings"

	conversions "github.com/GalacticDocs/Conversions"
	"github.com/iVitaliya/logger-go"
)

func EqualsTo(statement_one interface{}, statement_two interface{}) bool {
	return statement_one == statement_two
}

type IContains struct {
	Byte       func([]byte) bool
	Complex64  func(complex64) bool
	Complex128 func(complex128) bool
	Float32    func(float32) bool
	Float64    func(float64) bool
	Int        func(int) bool
	Int8       func(int8) bool
	Int16      func(int16) bool
	Int32      func(int32) bool
	Int64      func(int64) bool
	Rune       func(rune) bool
	String     func(string) bool
	UInt       func(uint) bool
	UIntPtr    func(uintptr) bool
	UInt8      func(uint8) bool
	UInt16     func(uint16) bool
	UInt32     func(uint32) bool
	UInt64     func(uint64) bool
}

func Contains(original interface{}) *IContains {
	FatalLog := func(type_def string) {
		L := logger.ConfigureLogger(logger.LoggerConfig{Severity: logger.FatalState})
		L.Log("The parameter \"original\" must be equal to a type of " + type_def + "!")
	}

	return &IContains{
		Byte: func(value []byte) bool {
			var (
				search []byte = value
				i, ok         = original.([]byte)
			)

			if !ok {
				FatalLog("[]byte")
			}

			return bytes.Contains(i, search)
		},
		Complex64: func(value complex64) bool {
			var (
				search string = conversions.Stringify().Complex64ToString(value)
				i, ok         = original.(complex64)
			)

			if !ok {
				FatalLog("complex64")
			}

			return strings.Contains(conversions.Stringify().Complex64ToString(i), search)
		},
		Complex128: func(value complex128) bool {
			var (
				search string = conversions.Stringify().Complex128ToString(value)
				i, ok         = original.(complex128)
			)

			if !ok {
				FatalLog("complex128")
			}

			return strings.Contains(conversions.Stringify().Complex128ToString(i), search)
		},
		Float32: func(value float32) bool {
			var (
				search string = conversions.Stringify().Float32ToString(value)
				i, ok         = original.(float32)
			)

			if !ok {
				FatalLog("float32")
			}

			return strings.Contains(conversions.Stringify().Float32ToString(i), search)
		},
		Float64: func(value float64) bool {
			var (
				search string = conversions.Stringify().Float64ToString(value)
				i, ok         = original.(float64)
			)

			if !ok {
				FatalLog("float64")
			}

			return strings.Contains(conversions.Stringify().Float64ToString(i), search)
		},
		Int: func(value int) bool {
			var (
				search string = conversions.Stringify().IntToString(value)
				i, ok         = original.(int)
			)

			if !ok {
				FatalLog("int")
			}

			return strings.Contains(conversions.Stringify().IntToString(i), search)
		},
		Int8: func(value int8) bool {
			var (
				search string = conversions.Stringify().Int8ToString(value)
				i, ok         = original.(int8)
			)

			if !ok {
				FatalLog("int8")
			}

			return strings.Contains(conversions.Stringify().Int8ToString(i), search)
		},
		Int16: func(value int16) bool {
			var (
				search string = conversions.Stringify().Int16ToString(value)
				i, ok         = original.(int16)
			)

			if !ok {
				FatalLog("int16")
			}

			return strings.Contains(conversions.Stringify().Int16ToString(i), search)
		},
		Int32: func(value int32) bool {
			var (
				search string = conversions.Stringify().Int32ToString(value)
				i, ok         = original.(int32)
			)

			if !ok {
				FatalLog("int32")
			}

			return strings.Contains(conversions.Stringify().Int32ToString(i), search)
		},
		Int64: func(value int64) bool {
			var (
				search string = conversions.Stringify().Int64ToString(value)
				i, ok         = original.(int64)
			)

			if !ok {
				FatalLog("int64")
			}

			return strings.Contains(conversions.Stringify().Int64ToString(i), search)
		},
		Rune: func(value rune) bool {
			var (
				search string = conversions.Stringify().RuneToString(value)
				i, ok         = original.(rune)
			)

			if !ok {
				FatalLog("rune")
			}

			return strings.Contains(conversions.Stringify().RuneToString(i), search)
		},
		String: func(value string) bool {
			var (
				search = value
				i, ok  = original.(string)
			)

			if !ok {
				FatalLog("string")
			}

			return strings.Contains(i, search)
		},
		UInt: func(value uint) bool {
			var (
				search = conversions.Stringify().UIntToString(value)
				i, ok  = original.(uint)
			)

			if !ok {
				FatalLog("uint")
			}

			return strings.Contains(conversions.Stringify().UIntToString(i), search)
		},
		UIntPtr: func(value uintptr) bool {
			var (
				search = conversions.Stringify().UIntPtrToString(value)
				i, ok  = original.(uintptr)
			)

			if !ok {
				FatalLog("uintptr")
			}

			return strings.Contains(conversions.Stringify().UIntPtrToString(i), search)
		},
		UInt8: func(value uint8) bool {
			var (
				search = conversions.Stringify().UInt8ToString(value)
				i, ok  = original.(uint8)
			)

			if !ok {
				FatalLog("uint8")
			}

			return strings.Contains(conversions.Stringify().UInt8ToString(i), search)
		},
		UInt16: func(value uint16) bool {
			var (
				search = conversions.Stringify().UInt16ToString(value)
				i, ok  = original.(uint16)
			)

			if !ok {
				FatalLog("uint16")
			}

			return strings.Contains(conversions.Stringify().UInt16ToString(i), search)
		},
		UInt32: func(value uint32) bool {
			var (
				search = conversions.Stringify().UInt32ToString(value)
				i, ok  = original.(uint32)
			)

			if !ok {
				FatalLog("uint32")
			}

			return strings.Contains(conversions.Stringify().UInt32ToString(i), search)
		},
		UInt64: func(value uint64) bool {
			var (
				search = conversions.Stringify().UInt64ToString(value)
				i, ok  = original.(uint64)
			)

			if !ok {
				FatalLog("uint64")
			}

			return strings.Contains(conversions.Stringify().UInt64ToString(i), search)
		},
	}
}
