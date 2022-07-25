package converter

import (
	"fmt"
	"strconv"
)

type IStringify struct {
	Complex64ToString  func(complex64) string
	Complex128ToString func(complex128) string
	Float32ToString    func(float32) string
	Float64ToString    func(float64) string
	IntToString        func(int) string
	Int8ToString       func(int8) string
	Int16ToString      func(int16) string
	Int32ToString      func(int32) string
	Int64ToString      func(int64) string
	RuneToString       func(rune) string
	UIntToString       func(uint) string
	UIntPtrToString    func(uintptr) string
	UInt8ToString      func(uint8) string
	UInt16ToString     func(uint16) string
	UInt32ToString     func(uint32) string
	UInt64ToString     func(uint64) string
}

func Stringify() *IStringify {
	return &IStringify{
		Complex64ToString: func(c complex64) string {
			var val string = fmt.Sprint(c)

			return val
		},
		Complex128ToString: func(c complex128) string {
			var val string = fmt.Sprint(c)

			return val
		},
		Float32ToString: func(f float32) string {
			var val string = fmt.Sprint(f)

			return val
		},
		Float64ToString: func(f float64) string {
			var val string = fmt.Sprint(f)

			return val
		},
		IntToString: func(i int) string {
			var val string = strconv.Itoa(i)

			return val
		},
		Int8ToString: func(i int8) string {
			var val string = strconv.Itoa(int(i))

			return val
		},
		Int16ToString: func(i int16) string {
			var val string = strconv.Itoa(int(i))

			return val
		},
		Int32ToString: func(i int32) string {
			var val string = strconv.Itoa(int(i))

			return val
		},
		Int64ToString: func(i int64) string {
			var val string = strconv.Itoa(int(i))

			return val
		},
		RuneToString: func(r rune) string {
			var val string = strconv.QuoteRune(r)

			return val
		},
		UIntToString: func(u uint) string {
			var val string = strconv.FormatUint(uint64(u), 10)

			return val
		},
		UIntPtrToString: func(u uintptr) string {
			var val string = strconv.FormatUint(uint64(u), 10)

			return val
		},
		UInt8ToString: func(u uint8) string {
			var val string = strconv.FormatUint(uint64(u), 10)

			return val
		},
		UInt16ToString: func(u uint16) string {
			var val string = strconv.FormatUint(uint64(u), 10)

			return val
		},
		UInt32ToString: func(u uint32) string {
			var val string = strconv.FormatUint(uint64(u), 10)

			return val
		},
		UInt64ToString: func(u uint64) string {
			var val string = strconv.FormatUint(uint64(u), 10)

			return val
		},
	}
}
