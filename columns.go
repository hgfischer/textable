package textable

import (
	"fmt"
)

type Column interface {
}

type StringerCol struct {
	Name string
	Rows []fmt.Stringer
}

type StringCol struct {
	Name string
	Rows []string
}

type UInt8Col struct {
	Name string
	Rows []uint8
}

type UInt16Col struct {
	Name string
	Rows []uint16
}

type UInt32Col struct {
	Name string
	Rows []uint32
}

type UInt64Col struct {
	Name string
	Rows []uint64
}

type Int8Col struct {
	Name string
	Rows []uint8
}

type Int16Col struct {
	Name string
	Rows []uint16
}

type Int32Col struct {
	Name string
	Rows []uint32
}

type Int64Col struct {
	Name string
	Rows []uint64
}

type Float32Col struct {
	Name string
	Rows []float64
}

type Float64Col struct {
	Name string
	Rows []float64
}

type Complex64Col struct {
	Name string
	Rows []complex64
}

type Complex128Col struct {
	Name string
	Rows []complex128
}

type BoolCol struct {
	Name string
	Rows []bool
}

type ByteCol struct {
	Name string
	Rows []byte
}

type RuneCol struct {
	Name string
	Rows []rune
}

type UIntCol struct {
	Name string
	Rows []uint
}

type IntCol struct {
	Name string
	Rows []uint
}
